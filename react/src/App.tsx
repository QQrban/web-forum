import { useEffect } from "react";
import { useState } from "react";
import "./App.css";
import { Header } from "./components/header";
import { Overlay } from "./components/shared/Overlay";
import { Banner } from "./components/banner";
import { Search } from "./components/search";
import { Home } from "./components/home";
import styled from "styled-components";
import { useDispatch } from "react-redux";
import { checkSession, getHomeData } from "./api/apiServices";
import { setLoginState } from "./redux/authSlice";
import { setHomePageData } from "./redux/mainDataSlice";

function App() {
  const [openOverlay, setOpenOverlay] = useState<boolean>(false);
  const dispatch = useDispatch();
  useEffect(() => {
    checkSession().then((data) => {
      data.Ok
        ? dispatch(setLoginState)
        : console.log("Guest Session is Active");
    });
    getHomeData().then((data) => {
      dispatch(setHomePageData(data));
    });
  }, [dispatch]);

  return (
    <>
      {openOverlay && <Overlay />}
      <Header setOpenOverlay={setOpenOverlay} />
      <Banner />
      <Search />
      <StyledMain>
        <Home />
      </StyledMain>
    </>
  );
}

export default App;

const StyledMain = styled.main`
  margin-top: 23px;
`;
