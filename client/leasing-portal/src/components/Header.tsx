import { Link } from "react-router-dom";
import { Button, Modal, ModalBody, ModalFooter, ModalHeader } from "reactstrap";
import React from "react";
import { useSelector, useDispatch } from "react-redux";
import Register from "./auth/Register";
import Login from "./auth/Login";
import {logoutSuccessful} from '../store/auth';
import  { useNavigate } from 'react-router-dom'

function Header() {
  // To-Do: toggle this
  const isLoggedIn = useSelector((state: any) => state.isAuthenticated.value);
 
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const logout = () => {
    dispatch(logoutSuccessful());
    navigate("./", {replace: true});
  }

  return (
    <div>
      <nav
        className="navbar navbar-expand-lg navbar-light ftco_navbar ftco-navbar-light"
        id="ftco-navbar"
      >
        <div className="container">
          <h1 className="navbar-brand">
            <strong>HOUS.</strong>
          </h1>
          <button
            className="navbar-toggler"
            type="button"
            data-toggle="collapse"
            data-target="#ftco-nav"
            aria-controls="ftco-nav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="oi oi-menu"></span> Menu
          </button>

          <div className="collapse navbar-collapse" id="ftco-nav">
            <ul className="navbar-nav ml-auto">
              <li className="nav-item active">
                <Link className="nav-link" to="/">
                  Home
                </Link>
              </li>

              {isLoggedIn ? (
                <>
                  <li className="nav-item">
                    <Link className="nav-link" to="/support">
                      Support
                    </Link>
                  </li>
                  <li className="nav-item">
                    <Link id="my-account" className="nav-link" to="/account">
                      My Account
                    </Link>
                  </li>
                  <li className="nav-item">
                    <span id="logout-button" className="nav-link" onClick={logout}>
                      Logout
                    </span>
                  </li>{" "}
                </>
              ) : (
                <>
                  <li className="nav-item">
                    <Login></Login>
                  </li>

                  <li className="nav-item">
                    <Register></Register>
                  </li>
                </>
              )}
            </ul>
          </div>
        </div>
      </nav>
    </div>
  );
}

export default Header;
