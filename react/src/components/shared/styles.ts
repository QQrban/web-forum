import styled from "styled-components";
import { variables } from "../../shared/variables";

export const bgWithStroke = styled.div`
  background: ${variables.mainBg};
  border: 2px solid ${variables.borderBlue};
  border-radius: 8px;
`;

export const buttonLeaf = styled.button`
  border-radius: 0px 16px;
  background: ${variables.lime};
  &:hover {
    background: ${variables.black};
    outline: 2px solid ${variables.lime};
    color: ${variables.lime};
  }
`;

export const StyledInput = styled.input`
  background: inherit;
  border-bottom: 1px solid ${variables.borderBlue};
  padding: 3px 10px;
  font-size: 20px;
  color: ${variables.white};
  font-family: ${variables.ffSquada};
`;

export const StyledForm = styled.form`
  display: flex;
  flex-direction: column;
  gap: 40px;
`;

export const InputsContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 20px;
`;
