import styled from "styled-components";
import {
  InputsContainer,
  StyledForm,
  StyledInput,
  bgWithStroke,
  buttonLeaf,
} from "../shared/styles";
import { variables } from "../../shared/variables";
import avatarExample from "../../assets/images/avatars/avocado.svg";
import { avatars } from "./avatars";
import { useState } from "react";
import { useOutsideClick } from "../../hooks/useOutsideHook";

type Props = {
  setOpenOverlay: React.Dispatch<React.SetStateAction<boolean>>;
  setOpenModal: React.Dispatch<React.SetStateAction<string>>;
};

export const Register = ({ setOpenOverlay, setOpenModal }: Props) => {
  const [avatarModal, setAvatarModal] = useState<boolean>(false);
  const [currentAvatar, setCurrentAvatar] = useState<string>(avatarExample);
  const ref = useOutsideClick(() => {
    setOpenModal("");
    setOpenOverlay(false);
  });

  const chooseAvatar = (avatar: string) => {
    setCurrentAvatar(avatar);
    setOpenOverlay(false);
    setAvatarModal(false);
  };

  const openAvatarModal = () => {
    setAvatarModal(true);
    setOpenOverlay(true);
  };

  return (
    <RegisterContainer ref={ref}>
      {avatarModal && (
        <AllAvatars>
          {avatars.map((avatar, index) => (
            <StyledImgAvatar
              onClick={() => chooseAvatar(avatar)}
              key={index}
              src={avatar}
              alt="avatar"
            />
          ))}
        </AllAvatars>
      )}
      <div>
        <LoginHeading>Sign Up</LoginHeading>
        <LoginSubHeading>
          Already a member?{" "}
          <SignUpButton onClick={() => setOpenModal("login")}>
            Log In
          </SignUpButton>
        </LoginSubHeading>
      </div>
      <AvatarsContainer>
        <AvatarsHeading>Choose Avatar</AvatarsHeading>
        <ImgContainer onClick={openAvatarModal}>
          <img src={currentAvatar} alt="avatar" />
        </ImgContainer>
      </AvatarsContainer>
      <RegisterForm>
        <InputsContainer>
          <StyledInput placeholder="Email" type="email" />
          <StyledInput placeholder="Username" type="text" />
          <StyledInput placeholder="Password" type="password" />
          <StyledInput placeholder="Confirm Password" type="password" />
        </InputsContainer>
        <StyledButton>Sign Up</StyledButton>
      </RegisterForm>
    </RegisterContainer>
  );
};

const RegisterContainer = styled(bgWithStroke)`
  position: absolute;
  width: 320px;
  right: 140px;
  top: 30px;
  border: 4px solid ${variables.borderBlue};
  padding: 35px;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-between;
  gap: 20px;
`;

const LoginHeading = styled.h3`
  font-size: 40px;
  text-align: center;
`;

const LoginSubHeading = styled.div`
  font-size: 19px;
  text-align: center;
  margin-top: 5px;
`;

const SignUpButton = styled.button`
  color: ${variables.lime};
  font-size: 19px;
  &:hover {
    text-decoration: underline;
  }
`;

const StyledButton = styled(buttonLeaf)`
  font-size: 32px;
  padding: 8px 16px;
`;

const AvatarsContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
`;

const ImgContainer = styled.div`
  width: 70px;
  height: 70px;
  overflow: hidden;
  border: 3px solid ${variables.white};
  border-radius: 50%;
  cursor: pointer;
  &:hover {
    opacity: 0.7;
  }
`;

const AvatarsHeading = styled.div`
  font-size: 19px;
`;

const AllAvatars = styled.div`
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  border: 4px solid ${variables.borderBlue};
  padding: 12px;
  display: flex;
  flex-wrap: wrap;
  background: ${variables.mainBg};
  z-index: 9999;
  gap: 10px;
`;

const StyledImgAvatar = styled.img`
  width: 60px;
  cursor: pointer;
  &:hover {
    outline: 1px solid ${variables.white};
  }
`;

const RegisterForm = styled(StyledForm)`
  align-items: center;
`;
