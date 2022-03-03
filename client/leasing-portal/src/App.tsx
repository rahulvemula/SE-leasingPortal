import "./App.css";
import Header from "./components/Header";
import Hero from "./components/Hero";
import Listings from "./components/Listings";
import Properties from "./components/Properties";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Footer from "./components/Footer";
import ApptListing from "./pages/ApptListing";
import Done from "./pages/done";
import { store } from "./store";
import { Provider } from "react-redux";
import "bootstrap/dist/css/bootstrap.min.css";
//import Account from "./pages/Account";

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
            <Route path="/done" element={<Done />} />
            {/* <Route path="/account" element = {<Account/>} /> */}
          </Routes>
          <Footer />
        </BrowserRouter>
      </div>
    </Provider>
  );
}

export default App;
