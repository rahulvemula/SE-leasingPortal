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
import AboutUs from "./components/AboutUs/AboutUs";

import Upload from "./components/UploadImage/UploadImage";
import Account from "./pages/Accounts";
import Listing from "./pages/Listing";
import Support from "./pages/Support";
import Terms from "./pages/Terms";

import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  const Application = () => {
    return (
      <>
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
              path="/home"
              element={
                <>
                  <Hero />
                  <Properties />
                  <Listings />
                </>
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

            <Route path="/support" element={<Support />} />
            <Route path="/appts" element={<ApptListing />} />
            <Route path="/lease-confirmation" element={<LeaseConfirmation />} />
            <Route path="/aboutus" element={<AboutUs />} />
            <Route path="/upload" element={<Upload/>} />
            {/* <Route path="/account" element = {<Account/>} /> */}
            <Route path="/account" element = {<Account/>} />
            <Route path="/listing/:id" element = {<Listing/>} />
            <Route path="/terms" element = {<Terms/>} />
          </Routes>
          <Footer />
      </>
    );
  };
  return (
    <Provider store={store}>
      <div className="App">
        <BrowserRouter basename="/SE-leasingPortal">
          <Application/>
        </BrowserRouter>
        <BrowserRouter>
          <Application/>
        </BrowserRouter>
      </div>
    </Provider>
  );
}

export default App;
