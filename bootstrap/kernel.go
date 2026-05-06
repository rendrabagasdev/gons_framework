package bootstrap

import (
	_ "gons/internal/cache"
	_ "gons/internal/database"
	_ "gons/internal/mailer"
	_ "gons/internal/queue"
	_ "gons/internal/service"
	_ "gons/internal/storage"
)
