package services

// ServiceProvider didefinisikan sebagai fungsi yang melakukan registrasi ke container
type ServiceProvider func() error

var ServiceRegistry []ServiceProvider

// RegisterService akan dipanggil oleh init() di setiap file service
func RegisterService(provider ServiceProvider) {
	ServiceRegistry = append(ServiceRegistry, provider)
}

// AutoRegisterServices akan dipanggil oleh config/services.go
func AutoRegisterService() {
	for _, register := range ServiceRegistry {
		register()
	}
}
