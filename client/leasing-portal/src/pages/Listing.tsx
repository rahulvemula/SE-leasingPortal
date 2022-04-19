import axios from "axios";
import { useState } from "react";
import { useDispatch, useSelector} from "react-redux";
import { useNavigate, useParams } from "react-router-dom";
import PropertyCard from "../components/Property-Card";
import {updateLeaseData} from '../store/leaseData';
import {Link} from 'react-router-dom';

function Listing() {
  let testData = {
    label: "Full house - 4bhk",
    rent: 2509,
    sqft: 2400,
    img: "House1",
    id: 1,
    isListingPage: true
  };

  const defaultData = {
    email: "",
    name: "",
    startDate: "",
    endDate: "",
  };

  const navigate = useNavigate();
  const dispatch = useDispatch();
  const storeLeaseData = useSelector((state: any) => state.leaseData);
  const [leaseData, setLeaseData] = useState(storeLeaseData);


  const updateName = (name: string) => {
    setLeaseData({ ...leaseData, name });
    dispatch(updateLeaseData(leaseData)); 
  };
  const updateEmail = (email: string) => {
    setLeaseData({ ...leaseData, email });
    dispatch(updateLeaseData(leaseData)); 
  };
  const updatestartDate = (startDate: string) => {
    setLeaseData({ ...leaseData, startDate });
    dispatch(updateLeaseData({ ...leaseData, startDate })); 
  };
  const updateEndDate = (endDate: string) => {
    setLeaseData({ ...leaseData, endDate });
    dispatch(updateLeaseData({ ...leaseData, endDate })); 
  };

  const [isChecked, setIsChecked] = useState(false);

  const handleOnChange = () => {
    setIsChecked(!isChecked);
  };

  const createLease = () => {
    if(!isChecked){
      alert("Please agree to terms and conditions!")
    } else{
      dispatch(updateLeaseData(leaseData));
      navigate('/lease-confirmation');
    }
    // axios
    // .post(
      // `${process.env.REACT_APP_API_URL}/users/${userData.username}`, leaseData, {headers : {token: res.data.token}}
      // )
    //   .then((response)=>{
    //     if(response.status == 200){
    //       navigate('/done');
    //     }
    //   })
    //   .finally(() => {
    //     setLeaseData(defaultData);
    //   });
  };

  let { id } = useParams();

  return (
    <section className="ftco-section goto-here">
      <div className="container">
        <div className="row" style={{margin:"15px"}}>
          <div className="col-md-1"></div>
          <div className="col-md-5">
            <div id="register-form">
              <div className="form-group">
                <label>Name</label>
                <input
                  type="text"
                  id="registerInputName"
                  className="form-control"
                  aria-describedby="nameHelp"
                  placeholder="John"
                  value={leaseData.name}
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
                  value={leaseData.email}
                  onChange={(e) => {
                    updateEmail(e.target.value);
                  }}
                />
              </div>
              <div className="form-group">
                <label>Start date</label>
                <input
                  id="registerInputStartDate"
                  type="date"
                  className="form-control"
                  placeholder="Password"
                  value={leaseData.startDate}
                  onChange={(e) => {
                    updatestartDate(e.target.value);
                  }}
                />
              </div>
              <div className="form-group">
                <label>End date</label>
                <input
                  id="registerInputEndDate"
                  type="date"
                  className="form-control"
                  placeholder="Password"
                  value={leaseData.endDate}
                  onChange={(e) => {
                    updateEndDate(e.target.value);
                  }}
                />
              </div>
              <div className="topping">
                <input
                  type="checkbox"
                  id="terms-checkbox"
                  name="topping"
                  value="Paneer"
                  checked={isChecked}
                  onChange={handleOnChange}
                  style={{marginTop:"15px"}}
                />
                 <span style={{marginLeft:"5px"}}>Agree to all &nbsp;
                 <Link id="terms" to={`/terms`}>
                  Terms and Conditions
                 </Link>
                 
                 </span>
              </div>
            </div>
            <div className="text-center">
              <button
                id="registerSubmit"
                className="btn btn-primary"
                onClick={createLease}
                style={{margin:"35px"}}
              >
                Generate Lease Document
              </button>
            </div>
          </div>
          <div className="col-md-1"></div>
          <div className="col-md-4">
            <PropertyCard {...testData}></PropertyCard>
          </div>
          <div className="col-md-1"></div>
        </div>
      </div>
    </section>
  );
}

export default Listing;
