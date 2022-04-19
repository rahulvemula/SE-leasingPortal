function Listings() {
  return (
    <section className="ftco-section goto-here">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-md-12 heading-section text-center">
            <span className="subheading">What we offer</span>
            <h2 className="mb-2">Exclusive Offer For You</h2>
          </div>
        </div>
        <br />
        <div className="row">
          <div className="col-md-4">
            <div className="property-wrap ">
              <div
                className="img d-flex align-items-center justify-content-center"
                style={{
                  backgroundImage: "url(SE-leasingPortal/images/House1.webp)",
                }}
              >
                <a
                  href="properties-single.html"
                  className="icon d-flex align-items-center justify-content-center btn-custom"
                >
                  <span className="ion-ios-link"></span>
                </a>
                <div className="list-agent d-flex align-items-center">
                  <a href="#" className="agent-info d-flex align-items-center">
                    <div
                      className="img-2 rounded-circle"
                      style={{
                        backgroundImage:
                          "url(SE-leasingPortal/images/xperson_1.jpg.pagespeed.ic.a2MnMHMs44.webp)",
                      }}
                    ></div>
                    <h3 className="mb-0 ml-2">Ben Ford</h3>
                  </a>
                  <div className="tooltip-wrap d-flex">
                    <a
                      href="#"
                      className="icon-2 d-flex align-items-center justify-content-center"
                      data-toggle="tooltip"
                      data-placement="top"
                      title=""
                      data-original-title="Bookmark"
                    >
                      <span className="ion-ios-heart">
                        <i className="sr-only">Bookmark</i>
                      </span>
                    </a>
                    <a
                      href="#"
                      className="icon-2 d-flex align-items-center justify-content-center"
                      data-toggle="tooltip"
                      data-placement="top"
                      title=""
                      data-original-title="Compare"
                    >
                      <span className="ion-ios-eye">
                        <i className="sr-only">Compare</i>
                      </span>
                    </a>
                  </div>
                </div>
              </div>
              <div className="text">
                <p className="price mb-3">
                  <span className="old-price">800,000</span>
                  <span className="orig-price">
                    $3,050<small>/mo</small>
                  </span>
                </p>
                <h3 className="mb-0">
                  <a href="properties-single.html">Blue View Home</a>
                </h3>
                <span className="location d-inline-block mb-3">
                  <i className="ion-ios-pin mr-2"></i>2854 Meadow View Drive,
                  Hartford, USA
                </span>
                <ul className="property_list">
                  <li>
                    <span className="flaticon-bed"></span>3
                  </li>
                  <li>
                    <span className="flaticon-bathtub"></span>2
                  </li>
                  <li>
                    <span className="flaticon-floor-plan"></span>1,878 sqft
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default Listings;
