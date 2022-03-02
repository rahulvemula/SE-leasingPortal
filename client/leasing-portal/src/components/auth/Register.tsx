import React, { useState } from "react";
import { Modal, ModalBody } from "reactstrap";
import axios from "axios";

function Register() {
  const [modal, setModal] = React.useState(false);
  const defaultData = {
    email: "",
    name: "",
    password: "",
  };
  const [userData, setUserData] = useState(defaultData);

  const updateName = (name: string) => {
    setUserData({ ...userData, name });
  };
  const updateEmail = (email: string) => {
    setUserData({ ...userData, email });
  };
  const updatePassword = (password: string) => {
    setUserData({ ...userData, password });
  };
  const register = () => {
    axios.post("https://murmuring-earth-87031.herokuapp.com/users", userData).finally(() => {
      setUserData(defaultData);
      setModal(false);
    });
  };

  // Toggle for Modal
  const toggle = () => setModal(!modal);

  return (
    <>
      <span id="register" className="nav-link" onClick={toggle}>
        Register
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
            <div id="register-form">
              <div className="form-group">
                <label>Name</label>
                <input
                  type="text"
                  id="registerInputName"
                  className="form-control"
                  aria-describedby="nameHelp"
                  placeholder="John"
                  value={userData.name}
                  onChange={(e) => {
                    updateName(e.target.value);
                  }}
                />
              </div>
              <div className="form-group">
                <label>Email address</label>
                <input
                  id="registerInputEmail"
                  type="email"
                  className="form-control"
                  aria-describedby="emailHelp"
                  placeholder="johndoe@example.com"
                  value={userData.email}
                  onChange={(e) => {
                    updateEmail(e.target.value);
                  }}
                />
              </div>
              <div className="form-group">
                <label>Password</label>
                <input
                  id="registerInputPassword"
                  type="password"
                  className="form-control"
                  placeholder="Password"
                  value={userData.password}
                  onChange={(e) => {
                    updatePassword(e.target.value);
                  }}
                />
              </div>
            </div>
            <div className="text-center">
              <button
                id="registerSubmit"
                className="btn btn-primary"
                onClick={register}
              >
                Submit
              </button>
            </div>
          </ModalBody>
        </Modal>
      </div>
    </>
  );
}

export default Register;
