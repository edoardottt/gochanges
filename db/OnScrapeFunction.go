package db

func (dbc *DatabaseConnection) OnScrapeFunction(scrapeTarget ScrapeTarget) {
	dbc.InsertScrapeTargets([]ScrapeTarget{scrapeTarget})
}
