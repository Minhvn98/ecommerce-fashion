package services

import "github.com/Minhvn98/ecommerce-fashion/services/cron"

func RunService() {
	cron.SendEmailAboutRevenueOfTheDay()
	cron.C.Start()
}
