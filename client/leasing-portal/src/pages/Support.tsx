function Support() {
  const filterOptions = [
    "Bedroom",
    "Living room",
    "Patio",
    "Bathroom",
    "Kitchen",
    "Laundry Room",
    "Other"
  ];

  const priorityOptions = ["High", "Medium", "Low"];

  const registerComplaint = () => {
      window.alert("Complaint registered, we will get back to you!");
  }

  return (
    <>
      <section className="ftco-section goto-here">
        <div className="container">
          <div className="row" >
            <div className="col"></div>
            {/* <div className="col-md-2"></div> */}
            <div className="col-5">
              <div id="register-form">
                <div className="form-group">
                  <label>Name</label>
                  <input
                    type="text"
                    id="registerInputName"
                    className="form-control"
                    aria-describedby="nameHelp"
                    placeholder="John"
                  />
                </div>

                <div className="form-group">
                  <label>Leased Property</label>
                  <input
                    type="text"
                    id="registerInputProperty"
                    className="form-control"
                    aria-describedby="nameHelp"
                    placeholder="The Niche appts"
                  />
                </div>

                <div className="form-group">
                  <label>Issue Location</label>
                  <br />

                  <div className="search-wrap-1">
                    <form action="#" className="search-property-1">
                      <div className="row">
                        <div className="col-lg align-items-end">
                          <div className="form-group">
                            <div className="form-field">
                              <div className="select-wrap">
                                <div className="icon">
                                  <i className="bi bi-caret-down-fill"></i>
                                </div>
                                <select
                                  name="locations"
                                  id="locations"
                                  data-testid="locationsSelect"
                                  className="form-control"
                                >
                                  {filterOptions.map((a) => (
                                    <option key={a} value={a}>
                                      {a}
                                    </option>
                                  ))}
                                </select>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </form>
                  </div>
                </div>

                <div className="form-group">
                  <label>Priority</label>
                  <br />

                  <div className="search-wrap-1">
                    <form action="#" className="search-property-1">
                      <div className="row">
                        <div className="col-lg align-items-end">
                          <div className="form-group">
                            <div className="form-field">
                              <div className="select-wrap">
                                <div className="icon">
                                  <i className="bi bi-caret-down-fill"></i>
                                </div>
                                <select
                                  name="locations"
                                  id="priority"
                                  data-testid="locationsSelect"
                                  className="form-control"
                                >
                                  {priorityOptions.map((a) => (
                                    <option key={a} value={a}>
                                      {a}
                                    </option>
                                  ))}
                                </select>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </form>
                  </div>
                </div>

                <div className="form-group">
                  <label>Issue description</label>
                  <textarea
                    id="registerInputDescription"
                    className="form-control"
                    aria-describedby="nameHelp"
                  />
                </div>
              </div>
            </div>
            <div className="col"></div>
          </div>

          <div className="text-center">
              <button
                id="registerSubmit"
                className="btn btn-primary"
                onClick={registerComplaint}
                style={{margin:"35px"}}
              >
                Register Complaint
              </button>
            </div>
        </div>
      </section>
    </>
  );
}

export default Support;
