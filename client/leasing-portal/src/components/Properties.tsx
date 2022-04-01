import { Link } from "react-router-dom";
function Properties() {
  // const navigate = useNavigate();
  return (
    <section className="ftco-section" id="properties">
      <div className="container">
        <div className="row justify-content-center">
          <div className="col-md-12 heading-section text-center mb-5">
            <span className="subheading">Find Properties</span>
            <h2 className="mb-2">Find Properties In Your City</h2>
          </div>
        </div>
        <div className="row">
          <div className="col-md-4">
            <div
              className="listing-wrap img rounded d-flex align-items-end"
              style={{
                backgroundImage: "url(images/Property1.webp)",
              }}
            >
              <div className="location">
                <span className="rounded">Gainesville, FL</span>
              </div>
              <div className="text">
                <h3>
                  <Link to="/appts">10 Property Listing</Link>
                </h3>
                <Link to="/appts" className="btn-link">
                  <span>
                    See All Listings{" "}
                    <i className="bi bi-arrow-right-circle-fill"></i>
                  </span>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default Properties;
