// import { useEffect } from "react";
import { useState } from "react";
import "./App.css";
import { Header } from "./components/header";
import { Overlay } from "./components/shared/Overlay";
// import { useDispatch } from "react-redux";
// import { checkSession } from "./api/apiServices";
// import { setLoginState } from "./redux/authSlice";

function App() {
  const [openOverlay, setOpenOverlay] = useState<boolean>(false);
  // const dispatch = useDispatch();
  // useEffect(() => {
  //   checkSession().then((data) => {
  //     data.Ok
  //       ? dispatch(setLoginState)
  //       : console.log("Guest Session is Active");
  //   });
  // }, [dispatch]);

  return (
    <>
      {openOverlay && <Overlay />}
      <Header setOpenOverlay={setOpenOverlay} />
    </>
  );
}

export default App;
