import "./App.css";
import Header from "./components/Header";
import Hero from "./components/Hero";
import Listings from "./components/Listings";
import Properties from "./components/Properties";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Footer from "./components/Footer";
import ApptListing from "./pages/ApptListing";
import LeaseConfirmation from "./pages/lease-confirmation";
import { store } from "./store";
import { Provider } from "react-redux";
import AboutUs from './components/AboutUs/AboutUs';

//import Account from "./pages/Account";
import Upload from './components/UploadImage/UploadImage';
import Account from "./pages/Accounts";

//import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/css/bootstrap.min.css';
function App() {
  return (
    <Provider store={store}>
      <div className="App">
        <BrowserRouter>
          <Header />
          <Routes>
            <Route
              path="*"
              element={
                <section className="ftco-section">
                  <span>not found</span>
                </section>
              }
            />
            <Route
              path="/"
              element={
                <>
                  <Hero />
                  <Properties />
                  <Listings />
                </>
              }
            />
            {/* TO-DO: add route to support page  */}
            <Route path="/support" element={<ApptListing />} />
            <Route path="/appts" element={<ApptListing />} />
            <Route path="/lease-confirmation" element={<LeaseConfirmation />} />
            <Route path="/aboutus" element={<AboutUs />} />
            <Route path="/upload" element={<Upload/>} />
            {/* <Route path="/account" element = {<Account/>} /> */}
            <Route path="/account" element = {<Account/>} />
          </Routes>
          <Footer />
        </BrowserRouter>
      </div>
    </Provider>
  );
}

export default App;
