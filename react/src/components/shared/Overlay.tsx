import styled from "styled-components";

export const Overlay = () => {
  return <StyledOverlay></StyledOverlay>;
};

const StyledOverlay = styled.div`
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  border: 0;
  width: 100%;
  z-index: 9989;
  background: gray;
  opacity: 0.5;
`;
