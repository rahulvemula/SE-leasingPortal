import "./App.css";
import Header from "./components/Header";
import Hero from "./components/Hero";
import Listings from "./components/Listings";
import Properties from "./components/Properties";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Footer from "./components/Footer";
import ApptListing from './pages/ApptListing';
import Done from './pages/done';
import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
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
        </Routes>
        <Footer />
      </BrowserRouter>
    </div>
  );
}

export default App;
