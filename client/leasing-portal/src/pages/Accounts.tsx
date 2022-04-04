import { useSelector } from "react-redux";
import "../index.css";

function Account() {
  const userData = useSelector((state: any) => {
    return state.userData;
  });

  const editProfile = () => {
    console.log("clickable button");
  };

  return (
    <>
      <section className="ftco-section goto-here">
        <div className="container">
          <section className="vh-100">
            <div className="container py-5 h-100">
              <div className="row d-flex justify-content-center align-items-center h-100">
                <div className="col col-lg-6 mb-4 mb-lg-0">
                  <div className="card mb-3">
                    <div className="row g-0">
                      <div className="col-md-4 gradient-custom text-center">
                        <img
                          src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-chat/ava3-bg.webp"
                          alt="Avatar"
                          className="img-fluid my-5"
                          style={{ width: "80px" }}
                        />
                        <h3>{userData.name}</h3>
                        <h6>{userData.username}</h6>
                        <br />
                        <button className="btn" onClick={editProfile}>
                          <i className="bi bi-pencil-square"></i>
                        </button>
                      </div>
                      <div className="col-md-8">
                        <div className="card-body p-4">
                          <h6>Information</h6>
                          <hr className="mt-0 mb-4" />
                          <div className="row pt-1">
                            <div className="col-6 mb-3">
                              <h6 id="name">Name</h6>
                              <p id="info-name" className="text-muted">{userData.name}</p>
                            </div>
                          </div>

                          <div className="row pt-1">
                            <div className="col-6 mb-3">
                              <h6>Leased Property</h6>
                              <p className="text-muted">The Niche Student Appartments</p>
                            </div>
                          </div>

                          <h6>Contact Information</h6>
                          <hr className="mt-0 mb-4" />
                          <div className="row pt-1">
                            <div className="col-6 mb-3">
                              <h6 id="email">Email</h6>
                              <p id="info-email" className="text-muted">{userData.email}</p>
                            </div>
                            <div className="col-6 mb-3">
                              <h6 id="phone">Phone number</h6>
                              <p className="text-muted">xxxxxxxx</p>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </section>
    </>
  );
}

export default Account;
