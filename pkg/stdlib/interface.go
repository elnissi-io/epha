package stdlib

type Module interface {
    Initialize() error // Initialize any necessary module setup.
    Name() string      // Provides the module name.
}
