import logo from './logo.svg';
import './App.css';
import Navbar from './components/Navbar';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import ProductsPage from './pages/ProductsPage';

function App() {
  return (
    
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path="/" element={<ProductsPage />} />
        {/* <Route path="/historical-value" element={<HistoricallyAvailableValue />} /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
