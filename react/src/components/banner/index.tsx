import styled from "styled-components";
import { bgWithStroke } from "../shared/styles";
import bg from "../../assets/images/background.jpeg";

export const Banner = () => {
  return (
    <BannerContainer>
      <BannerHeader>Welcome to Kood/Boards</BannerHeader>
      <BannerSubHeader>
        A platform created by students for students! Join today!
      </BannerSubHeader>
    </BannerContainer>
  );
};

const BannerContainer = styled(bgWithStroke)`
  flex-shrink: 0;
  margin-top: 23px;
  height: 240px;
  background: url(${bg});
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  cursor: pointer;
`;

const BannerHeader = styled.h1`
  font-size: 56px;
  margin-top: 75px;
  width: 100%;
  padding: 10px;
  text-align: center;
  background: #000;
`;

const BannerSubHeader = styled.div`
  font-size: 30px;
  align-self: flex-end;
  padding: 10px;
`;
