package stdlib

import (
	"fmt"
	"reflect"
	"sync"
)

var (
	moduleRegistry = make(map[string]func() Module)
	resourceRegistry = make(map[string]reflect.Type)
	regLock        sync.Mutex
)

// RegisterModule registers a module by its name and a constructor function.
func RegisterModule(name string, constructor func() Module) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, exists := moduleRegistry[name]; exists {
		panic(fmt.Sprintf("module already registered: %s", name))
	}
	moduleRegistry[name] = constructor
}

type ImportContext struct {
	imports map[string]Module
}

func NewImportContext() *ImportContext {
	return &ImportContext{
		imports: make(map[string]Module),
	}
}

func (ctx *ImportContext) ImportModule(name, alias string) error {
	moduleConstructor, ok := moduleRegistry[name]
	if !ok {
		return fmt.Errorf("module not found: %s", name)
	}
	module := moduleConstructor()
	if err := module.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize module '%s': %w", name, err)
	}
	ctx.imports[alias] = module
	return nil
}

func (ctx *ImportContext) ImportResourceType(name string, objType reflect.Type) {
	regLock.Lock()
	defer regLock.Unlock()
	resourceRegistry[name] = objType
}

func GetResourceType(name string) (reflect.Type, error) {
	regLock.Lock()
	defer regLock.Unlock()
	objType, exists := resourceRegistry[name]
	if !exists {
		return nil, fmt.Errorf("resource not found: %s", name)
	}
	return objType, nil
}
