import './App.css';
import Header from './components/Header';
import Hero from './components/Hero';
import Filters from './components/Filters';
import Listings from "./components/Listings";
import Properties from "./components/Properties";
import { BrowserRouter, Routes, Route } from "react-router-dom";
//import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
      <Header/>
        <Routes>
          <Route path="/" element={<><Hero /><Properties/><Listings/></>} />
          {/* TO-DO: add route to support page  */}
          <Route path="/support" element={<Filters />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
