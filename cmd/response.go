package cmd

type Response struct {
	Result Result `json:"result"`
}

type Result struct {
	Input          Input          `json:"input"`
	AddressMatches []AddressMatch `json:"addressMatches"`
}

type AddressMatch struct {
	MatchedAddress    string            `json:"matchedAddress"`
	Coordinates       Coordinates       `json:"coordinates"`
	TigerLine         TigerLine         `json:"tigerLine"`
	AddressComponents map[string]string `json:"addressComponents"`
	Geographies       Geographies       `json:"geographies"`
}

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Geographies struct {
	CensusBlockGroups []CensusBlockGroup `json:"Census Block Groups"`
	CensusBlocks      []CensusBlock      `json:"Census Blocks"`
}

type CensusBlockGroup struct {
	Pop100    string `json:"POP100"`
	Geoid     string `json:"GEOID"`
	Centlat   string `json:"CENTLAT"`
	Areawater int64  `json:"AREAWATER"`
	State     string `json:"STATE"`
	Basename  string `json:"BASENAME"`
	OID       int64  `json:"OID"`
	Lsadc     string `json:"LSADC"`
	Funcstat  string `json:"FUNCSTAT"`
	Intptlat  string `json:"INTPTLAT"`
	Name      string `json:"NAME"`
	Objectid  int64  `json:"OBJECTID"`
	Tract     string `json:"TRACT"`
	Centlon   string `json:"CENTLON"`
	Blkgrp    string `json:"BLKGRP"`
	Hu100     string `json:"HU100"`
	Arealand  int64  `json:"AREALAND"`
	Intptlon  string `json:"INTPTLON"`
	Mtfcc     string `json:"MTFCC"`
	Ur        string `json:"UR"`
	County    string `json:"COUNTY"`
}

type CensusBlock struct {
	Suffix    string `json:"SUFFIX"`
	Pop100    string `json:"POP100"`
	Geoid     string `json:"GEOID"`
	Centlat   string `json:"CENTLAT"`
	Block     string `json:"BLOCK"`
	Areawater int64  `json:"AREAWATER"`
	State     string `json:"STATE"`
	Basename  string `json:"BASENAME"`
	OID       int64  `json:"OID"`
	Lsadc     string `json:"LSADC"`
	Funcstat  string `json:"FUNCSTAT"`
	Intptlat  string `json:"INTPTLAT"`
	Name      string `json:"NAME"`
	Objectid  int64  `json:"OBJECTID"`
	Tract     string `json:"TRACT"`
	Centlon   string `json:"CENTLON"`
	Blkgrp    string `json:"BLKGRP"`
	Arealand  int64  `json:"AREALAND"`
	Hu100     string `json:"HU100"`
	Intptlon  string `json:"INTPTLON"`
	Mtfcc     string `json:"MTFCC"`
	Lwblktyp  string `json:"LWBLKTYP"`
	Ur        string `json:"UR"`
	County    string `json:"COUNTY"`
}

type TigerLine struct {
	TigerLineID string `json:"tigerLineId"`
	Side        string `json:"side"`
}

type Input struct {
	Benchmark Benchmark `json:"benchmark"`
	Vintage   Vintage   `json:"vintage"`
	Address   Address   `json:"address"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Benchmark struct {
	ID                   string `json:"id"`
	BenchmarkName        string `json:"benchmarkName"`
	BenchmarkDescription string `json:"benchmarkDescription"`
	IsDefault            bool   `json:"isDefault"`
}

type Vintage struct {
	ID                 string `json:"id"`
	VintageName        string `json:"vintageName"`
	VintageDescription string `json:"vintageDescription"`
	IsDefault          bool   `json:"isDefault"`
}
