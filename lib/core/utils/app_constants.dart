enum MarketStatus {
  MARKET_OPEN,
  MARKET_CLOSED,
  MARKET_CLOSING,
  MARKET_OPENING,
}

extension HeaderType on MarketStatus {
  String get value {
    switch (this) {
      case MarketStatus.MARKET_OPEN:
        return "lbl_market_open";
      case MarketStatus.MARKET_CLOSED:
        return "lbl_market_closed";
      case MarketStatus.MARKET_CLOSING:
        return "lbl_market_closing";
      case MarketStatus.MARKET_OPENING:
        return "lbl_market_opening";
      default:
        return "";
    }
  }
}

enum StockDataAmountType {
  PRICE,
  VOL_CR,
  VALUE,
  NEW_HIGH,
  NEW_LOW,
}

extension AmountType on StockDataAmountType {
  String get value {
    switch (this) {
      case StockDataAmountType.PRICE:
        return "lbl_price";
      case StockDataAmountType.VOL_CR:
        return "lbl_vol_cr";
      case StockDataAmountType.VALUE:
        return "lbl_value_cr";
      case StockDataAmountType.NEW_HIGH:
        return "lbl_new_high";
      case StockDataAmountType.NEW_LOW:
        return "lbl_new_low";
      default:
        return "";
    }
  }
}

enum SectionTabValue {
  WEEKHIGH_52,
  WEEKLOW_52,
  GAINERS,
  LOOSERS,
  BY_VALUE,
  BY_VOLUME
}

extension TabName on SectionTabValue {
  String get value {
    switch (this) {
      case SectionTabValue.GAINERS:
        return "lbl_gainers";
      case SectionTabValue.LOOSERS:
        return "lbl_losers";
      case SectionTabValue.BY_VALUE:
        return "lbl_by_value";
      case SectionTabValue.BY_VOLUME:
        return "lbl_by_volume";
      case SectionTabValue.WEEKHIGH_52:
        return "lbl_52_week_highs";
      case SectionTabValue.WEEKLOW_52:
        return "lbl_52_week_lows";

      default:
        return "";
    }
  }
}
