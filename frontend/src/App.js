import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { MainPage } from './pages/mainpage';
import {Webchat} from './pages/webchat';
import {SingUp} from './pages/signup';

const App = () => {
  return (
    <BrowserRouter> 
      <Routes>
        <Route path={"/"} element={<MainPage />}/>
        <Route path={"/webchat"} element={<Webchat />}/>
        <Route path={"/sing-up"} element={<SingUp />}/>
      </Routes>
    </BrowserRouter> 
  );
}

export default App;