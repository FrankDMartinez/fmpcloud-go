package fmpcloud

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPIStockSymbolList                = "/stock/list"
	urlAPIStockSymbolByExchangeList      = "/symbol/%s"
	urlAPIStockCompanyProfile            = "/profile/%s"
	urlAPIStockCompanyExecutives         = "/key-executives/%s"
	urlAPIStockDividends                 = "/historical-price-full/stock_dividend/%s"
	urlAPIStockSplits                    = "/historical-price-full/stock_split/%s"
	urlAPIStockQuote                     = "/quote/%s"
	urlAPIStockQuoteShot                 = "/quote-short/%s"
	urlAPIStockQuotes                    = "/quotes/%s"
	urlAPIStockSearch                    = "/search"
	urlAPIStockSearchTicker              = "/search-ticker"
	urlAPIStockCandles                   = "/historical-chart/%s/%s"
	urlAPIStockDaily                     = "/historical-price-full/%s"
	urlAPIStockSP500List                 = "/sp500_constituent"
	urlAPIStockHistorySP500List          = "/historical/sp500_constituent"
	urlAPIStockDowJonesList              = "/dowjones_constituent"
	urlAPIStockHistoryDowJonesList       = "/historical/dowjones_constituent"
	urlAPIStockNasdaqList                = "/nasdaq_constituent"
	urlAPIStockHistoryNasdaqList         = "/historical/nasdaq_constituent"
	urlAPIStockEODCandles                = "/batch-request-end-of-day-prices"
	urlAPIStockEODBatchCandles           = "/batch-request-end-of-day-prices/%s"
	urlAPIStockMarketHours               = "/market-hours"
	urlAPIStockActives                   = "/actives"
	urlAPIStockLosers                    = "/losers"
	urlAPIStockGainers                   = "/gainers"
	urlAPIStockSectorsPerformance        = "/sectors-performance"
	urlAPIStockHistorySectorsPerformance = "/historical-sectors-performance"
)

// Stock client
type Stock struct {
	Client *resty.Client
	url    string
	apiKey string
}

// QuoteShort - real-time single quote short
func (s *Stock) QuoteShort(symbol string) (qList []objects.StockQuoteShot, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockQuoteShot, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// Quote - real-time single quote
func (s *Stock) Quote(symbol string) (qList []objects.StockQuote, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockQuote, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// BatchQuote - real-time batch quote
func (s *Stock) BatchQuote(symbolList []string) (qList []objects.StockQuote, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockQuote, strings.Join(symbolList, ",")))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// QuoteByExchange - real-time single quote
func (s *Stock) QuoteByExchange(exchange objects.StockSearch) (qList []objects.StockQuote, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockQuotes, string(exchange)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// Search - ticker search exchange (nasdaq | nyse | tsx | euronext | mutual_fund | etf | amex | index | commodity | forex | crypto)
func (s *Stock) Search(req objects.RequestStockSearch) (sList []objects.StockSymbol, err error) {
	reqParam := map[string]string{
		"apikey": s.apiKey,
		"limit":  fmt.Sprint(req.Limit),
		"query":  req.Query,
	}
	if req.Exchange != nil {
		reqParam["exchange"] = req.Exchange.String()
	}

	data, err := s.Client.R().
		SetQueryParams(reqParam).
		Get(s.url + urlAPIStockSearch)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// SearchTiker - only ticker search exchange (nasdaq | nyse | tsx | euronext | mutual_fund | etf | amex | index | commodity | forex | crypto)
func (s *Stock) SearchTiker(req objects.RequestStockSearch) (sList []objects.StockSymbol, err error) {
	reqParam := map[string]string{
		"apikey": s.apiKey,
		"limit":  fmt.Sprint(req.Limit),
		"query":  req.Query,
	}
	if req.Exchange != nil {
		reqParam["exchange"] = req.Exchange.String()
	}

	data, err := s.Client.R().
		SetQueryParams(reqParam).
		Get(s.url + urlAPIStockSearchTicker)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// CompanyProfile - get general information of a company. You can query by symbol.
func (s *Stock) CompanyProfile(symbol string) (companyProfile []objects.StockCompanyProfile, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockCompanyProfile, symbol))

	if err != nil {
		return nil, err
	}

	if data.StatusCode() != http.StatusOK {
		return nil, errors.New(string(data.Body()))
	}

	err = json.Unmarshal(data.Body(), &companyProfile)
	if err != nil {
		return nil, err
	}

	return companyProfile, nil
}

// CompanyExecutive - get a list of company's executives and members of the Board.
func (s *Stock) CompanyExecutive(symbol string) (companyProfile []objects.CompanyExecutive, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockCompanyExecutives, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &companyProfile)
	if err != nil {
		return nil, err
	}

	return companyProfile, nil
}

// Candles - historical candles
func (s *Stock) Candles(req objects.RequestStockCandleList) (cList []objects.StockCandle, err error) {
	reqParam := map[string]string{"apikey": s.apiKey}
	if req.From != nil {
		reqParam["from"] = req.From.Format("2006-01-02")
	}

	if req.To != nil {
		reqParam["to"] = req.To.Format("2006-01-02")
	}

	data, err := s.Client.R().
		SetQueryParams(reqParam).
		Get(s.url + fmt.Sprintf(urlAPIStockCandles, req.Period, req.Symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLine - daily line
func (s *Stock) DailyLine(symbol string, serieType objects.StockSerieType) (cList *objects.StockDailyLineList, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey, "serietype": string(serieType)}).
		Get(s.url + fmt.Sprintf(urlAPIStockDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyChangeAndVolume - daily candle change and volume
func (s *Stock) DailyChangeAndVolume(symbol string) (cList *objects.StockDailyCandleList, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailySpecificPeriod - daily candle list by specific period
func (s *Stock) DailySpecificPeriod(symbol string, from time.Time, to time.Time) (cList *objects.StockDailyCandleList, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{
			"apikey": s.apiKey,
			"from":   from.Format("2006-01-02"),
			"to":     to.Format("2006-01-02"),
		}).
		Get(s.url + fmt.Sprintf(urlAPIStockDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLastNDays - daily candle list last N days
func (s *Stock) DailyLastNDays(symbol string, days int) (cList *objects.StockDailyCandleList, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey, "timeseries": fmt.Sprint(days)}).
		Get(s.url + fmt.Sprintf(urlAPIStockDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyBatch - daily candle list
func (s *Stock) DailyBatch(symbolList []string, from *time.Time, to *time.Time) (cList []objects.StockBatchData, err error) {
	reqParam := map[string]string{"apikey": s.apiKey}
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := s.Client.R().
		SetQueryParams(reqParam).
		Get(s.url + fmt.Sprintf(urlAPIStockDaily, strings.Join(symbolList, ",")))

	if err != nil {
		return nil, err
	}

	var resp objects.StockBatchDaily
	err = json.Unmarshal(data.Body(), &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// Dividends - stock dividends
func (s *Stock) Dividends(symbol string) (dList *objects.StockDividends, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockDividends, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &dList)
	if err != nil {
		return nil, err
	}

	return dList, nil
}

// Splits - stock splits
func (s *Stock) Splits(symbol string) (sList *objects.StockSplit, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockSplits, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// AvalibleSymbolsByExchange - symbol list by exchange
func (s *Stock) AvalibleSymbolsByExchange(exchange objects.StockSymbolExchange) (sList []objects.StockSymbol, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + fmt.Sprintf(urlAPIStockSymbolByExchangeList, string(exchange)))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// AvalibleSymbols - all avalible symbol list
func (s *Stock) AvalibleSymbols() (sList []objects.StockSymbolList, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockSymbolList)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// IndexConstituentList - list of index companies (SP500, Nasdaq, DJ)
func (s *Stock) IndexConstituentList(index objects.Index) (sList []objects.IndexSymbol, err error) {
	url := s.url
	switch index {
	case objects.IndexSP500:
		url += urlAPIStockSP500List
	case objects.IndexDowJones:
		url += urlAPIStockDowJonesList
	case objects.IndexNasdaq100:
		url += urlAPIStockNasdaqList
	}

	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(url)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// HistoryIndexConstituentList - historical index companies list (SP500, Nasdaq, DJ)
func (s *Stock) HistoryIndexConstituentList(index objects.Index) (sList []objects.HistoryIndexSymbol, err error) {
	url := s.url
	switch index {
	case objects.IndexSP500:
		url += urlAPIStockHistorySP500List
	case objects.IndexDowJones:
		url += urlAPIStockHistoryDowJonesList
	case objects.IndexNasdaq100:
		url += urlAPIStockHistoryNasdaqList
	}
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(url)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// EODCandleList - all stocks Batch EOD stock price
func (s *Stock) EODCandleList(date time.Time) (sList []objects.StockEODCandle, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{
			"apikey": s.apiKey,
			"date":   date.Format("2006-01-02"),
		}).
		Get(s.url + urlAPIStockEODCandles)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BatchEODCandleList - specific Stocks Batch EOD stock prices
func (s *Stock) BatchEODCandleList(symbolList []string, date time.Time) (sList []objects.StockEODCandle, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{
			"apikey": s.apiKey,
			"date":   date.Format("2006-01-02"),
		}).
		Get(s.url + fmt.Sprintf(urlAPIStockEODBatchCandles, strings.Join(symbolList, ",")))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// ExchangeTradingHours - stock market trading hours
func (s *Stock) ExchangeTradingHours() (eList []objects.Exchange, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockMarketHours)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// Actives - stock market top active
func (s *Stock) Actives() (aList []objects.Active, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockActives)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &aList)
	if err != nil {
		return nil, err
	}

	return aList, nil
}

// Losers - stock market top losers
func (s *Stock) Losers() (lList []objects.Loser, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockLosers)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &lList)
	if err != nil {
		return nil, err
	}

	return lList, nil
}

// Gainers - stock market top gainers
func (s *Stock) Gainers() (gList []objects.Gainer, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockGainers)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &gList)
	if err != nil {
		return nil, err
	}

	return gList, nil
}

// SectorPerformance - stock market sector performance
func (s *Stock) SectorPerformance() (eList []objects.Sector, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockSectorsPerformance)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// HistorySectorPerformance - historical stock market sector performance
func (s *Stock) HistorySectorPerformance() (eList []objects.HistorySector, err error) {
	data, err := s.Client.R().
		SetQueryParams(map[string]string{"apikey": s.apiKey}).
		Get(s.url + urlAPIStockHistorySectorsPerformance)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}
