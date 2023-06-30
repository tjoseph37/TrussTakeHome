package models
type Records struct{
	// Fields Timestamp,Address,ZIP,FullName,FooDuration,BarDuration,TotalDuration,Notes
	Timestamp string `csv:"Timestamp"`
	Address string `csv:"Address"`
	Zipcode string `csv:"ZIP"`
	FullName string `csv:"FullName"`
	FooDuration string `csv:"FooDuration"`
	BarDuration string `csv:"BarDuration"`
	TotalDuration string `csv:"TotalDuration"`
	Notes string `csv:"Notes"`
}