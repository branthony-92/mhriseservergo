import { Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import ArmourSetsPage from './pages/ArmourSetsPage';
import Layout from './components/layout/Layout';

function App() {
  return (
    <Layout>
      <Routes>
        <Route path='/' element={<HomePage/>} />
        <Route path='/armour-sets' element={<ArmourSetsPage/>} />
      </Routes>
    </Layout>
  );
}

export default App;
