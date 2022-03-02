import React, { useState } from "react";
import { Modal, ModalBody } from "reactstrap";
import axios from "axios";

function Login() {
  const [modal, setModal] = React.useState(false);
  const defaultData = {
    email: "",
    password: "",
  };
  const [userData, setUserData] = useState(defaultData);

  // Toggle for Modal
  const toggle = () => setModal(!modal);

  const authenticateUser = () => {};
  const updateEmail = (email: string) => {
    setUserData({ ...userData, email });
  };
  const updatePassword = (password: string) => {
    setUserData({ ...userData, password });
  };

  //TO-DO: Add login api and handle success and failure cases
  const login = () => {
    axios
      .post("http://localhost:9010/user/", userData)
      .then(() => {
        authenticateUser();
      })
      .finally(() => {
        setUserData(defaultData);
        setModal(false);
      });
  };

  return (
    <>
      <span className="nav-link" onClick={toggle}>
        Login
      </span>

      {/* Register  */}
      <div>
        <Modal
          centered
          isOpen={modal}
          toggle={toggle}
          modalTransition={{ timeout: 2 }}
        >
          <ModalBody>
            <div className="form-group">
              <label>Email address</label>
              <input
                type="email"
                className="form-control"
                id="loginInputEmail"
                aria-describedby="emailHelp"
                placeholder="johndoe@example.com"
                onChange={(e) => {
                  updateEmail(e.target.value);
                }}
              />
            </div>
            <div className="form-group">
              <label>Password</label>
              <input
                type="password"
                className="form-control"
                id="loginInputPassword"
                placeholder="Password"
                onChange={(e) => {
                  updatePassword(e.target.value);
                }}
              />
            </div>
            <div className="text-center">
              <button
                id="loginSubmit"
                className="btn btn-primary"
                onClick={login}
              >
                Login
              </button>
            </div>
          </ModalBody>
        </Modal>
      </div>
    </>
  );
}

export default Login;
