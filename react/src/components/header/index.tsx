import styled from "styled-components";
import logo from "../../assets/images/logo.png";
import { bgWithStroke, buttonLeaf } from "../shared/styles";
// import { useSelector } from "react-redux/es/hooks/useSelector";
// import { RootState } from "../../redux/store";
import myOffice from "../../assets/images/profile.svg";
import { Register } from "./Register";
import { useState } from "react";
import { Login } from "./Login";

type Props = {
  setOpenOverlay: React.Dispatch<React.SetStateAction<boolean>>;
};

export const Header = ({ setOpenOverlay }: Props) => {
  // const isLogged = useSelector((state: RootState) => state.auth.isLogged);
  // console.log(isLogged);
  const isLogged = false;
  const [openedModal, setOpenedModal] = useState<string>("");

  return (
    <HeaderStyled>
      <Logo>
        <LogoImgContainer>
          <img src={logo} alt="" />
        </LogoImgContainer>
        <LogoName>kood/boards</LogoName>
      </Logo>
      <ButtonsContainer>
        {!isLogged ? (
          <>
            <HeaderBtn onClick={() => setOpenedModal("register")}>
              Sign Up
            </HeaderBtn>
            <HeaderBtn onClick={() => setOpenedModal("login")}>
              Log In
            </HeaderBtn>
            {openedModal === "login" && <Login setOpenModal={setOpenedModal} />}
            {openedModal === "register" && (
              <Register
                setOpenModal={setOpenedModal}
                setOpenOverlay={setOpenOverlay}
              />
            )}
          </>
        ) : (
          <img src={myOffice} alt="" />
        )}
      </ButtonsContainer>
    </HeaderStyled>
  );
};

const HeaderStyled = styled(bgWithStroke)`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  color: white;
`;

const Logo = styled.div`
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
`;

const LogoImgContainer = styled.div`
  width: 40px;
  height: 40px;
  border-radius: 50%;
  overflow: hidden;
`;

const LogoName = styled.div`
  font-size: 30px;
  text-transform: uppercase;
`;

const ButtonsContainer = styled.div`
  display: flex;
  gap: 20px;
  align-items: center;
  position: relative;
`;

const HeaderBtn = styled(buttonLeaf)`
  padding: 5px 15px;
  font-size: 22px;
`;
