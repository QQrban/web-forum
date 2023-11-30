import styled from "styled-components";
import {
  InputsContainer,
  StyledForm,
  StyledInput,
  bgWithStroke,
  buttonLeaf,
} from "../shared/styles";
import { variables } from "../../shared/variables";
import { useOutsideClick } from "../../hooks/useOutsideHook";

type Props = {
  setOpenModal: React.Dispatch<React.SetStateAction<string>>;
};

export const Login = ({ setOpenModal }: Props) => {
  const ref = useOutsideClick(() => {
    setOpenModal("");
  });

  return (
    <LoginContainer ref={ref}>
      <div>
        <LoginHeading>Log In</LoginHeading>
        <LoginSubHeading>
          Not a member?{" "}
          <SignUpButton onClick={() => setOpenModal("register")}>
            Sign Up
          </SignUpButton>
        </LoginSubHeading>
      </div>
      <StyledForm>
        <InputsContainer>
          <StyledInput placeholder="Username" type="text" />
          <StyledInput placeholder="Password" type="password" />
        </InputsContainer>
        <StyledButton>Log In</StyledButton>
      </StyledForm>
    </LoginContainer>
  );
};

const LoginContainer = styled(bgWithStroke)`
  position: absolute;
  min-height: 350px;
  width: 320px;
  right: 30px;
  top: 30px;
  border: 4px solid ${variables.borderBlue};
  padding: 35px;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-between;
  z-index: 9999;
  gap: 30px;
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
  max-width: 150px;
  align-self: center;
  padding: 6px 20px;
`;
