import { useState } from "react";
import { Modal, ModalBody } from "reactstrap";
import axios from "axios";
import { useDispatch } from "react-redux";
import { loginSuccessful } from "../../store/auth";
import { updateUserData } from "../../store/userData";

function Login() {
  const [modal, setModal] = useState(false);
  const dispatch = useDispatch();
  const defaultData = {
    email: "",
    password: "",
  };
  const mockUserData = {
    name: "R",
    email: "R",
    username: "R"
  }
  const [userData, setUserData] = useState(defaultData);

 
  let showSpinner = false;

  // Toggle for Modal
  const toggle = () => setModal(!modal);

  const authenticateUser = () => {
    dispatch(loginSuccessful());
  };
  const updateEmail = (email: string) => {
    setUserData({ ...userData, email });
  };
  const updatePassword = (password: string) => {
    setUserData({ ...userData, password });
  };
  const updateStateWithUserData = (userData: any) => {
    dispatch(updateUserData(userData));
  }


  //TO-DO: Add login api and handle success and failure cases
  const login = () => {
    showSpinner = true;
    
    axios.get(`https://murmuring-earth-87031.herokuapp.com/users/${userData.email}`).then((res) => {
      if(res.data.ID != 0) {
        authenticateUser();
        const userResponse = {name: res.data.Name, email: res.data.Email, password: res.data.Password};
        updateStateWithUserData(userResponse);
      } else {
        // TO-DO: this needs to be handled in catch
        alert('Auth failed'); 
      }
    }).finally(() => {
      setModal(false);
      showSpinner = false;
    })



    // axios
    //   .post("http://localhost:9010/user/", userData)
    //   .then(() => {
    //     authenticateUser();
    //   })
    //   .finally(() => {
    //     setUserData(defaultData);
    //     setModal(false);
    //   });

    // authenticateUser();
    // updateStateWithUserData(mockUserData);
    // setUserData(defaultData);
    
  };

  return (
    <>
      <span id="login" className="nav-link" onClick={toggle}>
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
          <ModalBody id="login-form">
           
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
            <div className="text-center"> {
              showSpinner ?  
              <button className="btn btn-primary" type="button" disabled>
              <span className="spinner-grow spinner-grow-sm" role="status" aria-hidden="true"></span>
              &nbsp;Logging...
            </button>
              :
              <button
                id="loginSubmit"
                className="btn btn-primary"
                onClick={login}
              >
                Login
              </button>
              
            }
              
            </div>
         
          </ModalBody>
        </Modal>
      </div>
    </>
  );
}

export default Login;
