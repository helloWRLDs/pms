import { BrowserRouter } from "react-router-dom";
import Main from "./pages/Main.tsx";

function App() {
  return (
    <>
      <BrowserRouter>
        {/* <Header logoURL="./src/assets/logo.png" /> */}
        <Main />
      </BrowserRouter>
    </>
  );
}

export default App;
