import styled from "styled-components";
import { bgWithStroke } from "../shared/styles";
import { variables } from "../../shared/variables";
import { CiSearch } from "react-icons/ci";

export const Search = () => {
  return (
    <SearchContainer>
      <SearchInputContainer>
        <StyledSearchInput placeholder="Search..." type="text" />
        <StyledIcon />
      </SearchInputContainer>
      <SearchButtons>
        <SearchNavLinks href="/about">About</SearchNavLinks>
        <SearchNavLinks href="/">Home</SearchNavLinks>
      </SearchButtons>
    </SearchContainer>
  );
};

const SearchContainer = styled(bgWithStroke)`
  margin-top: 23px;
  padding: 15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const SearchInputContainer = styled.div`
  position: relative;
`;

const SearchButtons = styled.div`
  display: flex;
  gap: 20px;
`;

const SearchNavLinks = styled.a`
  font-size: 22px;
  &:hover {
    text-decoration: underline;
  }
`;

const StyledSearchInput = styled.input`
  width: 275px;
  height: 40px;
  padding: 10px;
  color: #222;
  font-size: 22px;
  border-radius: 8px;
  font-family: ${variables.ffSquada};
`;

const StyledIcon = styled(CiSearch)`
  position: absolute;
  right: 5px;
  color: ${variables.black};
  font-size: 30px;
  top: 50%;
  transform: translateY(-50%);
`;
