
const filterOptions = {
  locations: ["Gainesville", "Santa Fe", "Alachua"],
  type: [
    "Residential",
    "Commercial",
    "Office",
    "Villa",
    "Condominium",
    "Apartment",
  ],
  prices: [250, 500, 750, 1000, 2500, 5000],
};

function Filters() {
  return (
    <>
      <section className="ftco-section ftco-no-pb">
        <div className="container">
          <div className="row">
            <div className="col-md-12">
              <div className="search-wrap-1">
                <form action="#" className="search-property-1">
                  <div className="row">
                    <div className="col-lg align-items-end">
                      <div className="form-group">
                        <label htmlFor="#">Location</label>
                        <div className="form-field">
                          <div className="select-wrap">
                            <div className="icon">
                              <i className="bi bi-caret-down-fill"></i>
                            </div>
                            <select name="locations" id="locations" data-testid="locationsSelect" className="form-control">
                              {filterOptions.locations.map((a) => (
                                <option key={a} value={a}>
                                  {a}
                                </option>
                              ))}
                            </select>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div className="col-lg align-items-end">
                      <div className="form-group">
                        <label htmlFor="#">Property Type</label>
                        <div className="form-field">
                          <div className="select-wrap">
                            <div className="icon">
                              <i className="bi bi-caret-down-fill"></i>
                            </div>
                            <select name="type" id="ptype" data-testid="propTypesSelect" className="form-control">
                              {filterOptions.type.map((a) => (
                                <option key={a} value={a}>
                                  {a}
                                </option>
                              ))}
                            </select>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div className="col-lg align-items-end">
                      <div className="form-group">
                        <label htmlFor="#">Price Limit</label>
                        <div className="form-field">
                          <div className="select-wrap">
                            <div className="icon">
                              <i className="bi bi-caret-down-fill"></i>
                            </div>
                            <select name="prices" id="prices" data-testid="pricesSelect" className="form-control">
                              {filterOptions.prices.map((a) => (
                                <option key={a} value={a}>
                                  {a}
                                </option>
                              ))}
                            </select>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div className="col-lg align-self-end">
                      <div className="form-group">
                        <div className="form-field">
                          <input
                            type="submit"
                            value="Search Property"
                            className="form-control btn btn-primary"
                          />
                        </div>
                      </div>
                    </div>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}

export default Filters;
