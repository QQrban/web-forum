import {
  Accordion,
  AccordionItem,
  AccordionItemHeading,
  AccordionItemButton,
  AccordionItemPanel,
} from "react-accessible-accordion";

import { FaChevronDown } from "react-icons/fa";
import "react-accessible-accordion/dist/fancy-example.css";
import styled from "styled-components";
import { variables } from "../../shared/variables";
import topicImg from "../../assets/images/topicimg.png";
import { useSelector } from "react-redux";
import { RootState } from "../../redux/store";

export const Home = () => {
  const homeData = useSelector((state: RootState) => state.mainPage);

  return (
    <StyledAccordion
      preExpanded={["uuid1", "uuid2", "uuid3"]}
      allowMultipleExpanded={true}
      allowZeroExpanded={true}
    >
      {homeData
        ? homeData.categories.map((cat, index) => (
            <StyledAccordionItem
              key={cat.category.id}
              uuid={`uuid${index + 1}`}
            >
              <AccordionItemHeading>
                <StyledAccordionItemButton>
                  <div>
                    <AccordionHeading>{cat.category.title}</AccordionHeading>
                    <AccordionSubHeading>
                      {cat.category.intro}
                    </AccordionSubHeading>
                  </div>
                  <StyledChevron />
                </StyledAccordionItemButton>
              </AccordionItemHeading>
              {cat.topics.map((top) => (
                <StyledAccordionItemPanel>
                  <AccordionItemNameContainer>
                    <div>
                      <img src={topicImg} alt="topic" />
                    </div>
                    <div>
                      <TopicHeading>{top.category.title}</TopicHeading>
                      <TopicSubHeading>{top.category.intro}</TopicSubHeading>
                    </div>
                  </AccordionItemNameContainer>
                </StyledAccordionItemPanel>
              ))}
            </StyledAccordionItem>
          ))
        : "kek"}
    </StyledAccordion>
  );
};

const StyledAccordion = styled(Accordion)`
  display: flex;
  flex-direction: column;
  gap: 23px;
`;

const StyledAccordionItem = styled(AccordionItem)`
  border: 2px solid ${variables.borderBlue};
  border-radius: 8px;
  overflow: hidden;
`;

const StyledChevron = styled(FaChevronDown)`
  transition: transform 0.3s ease;
`;

const StyledAccordionItemButton = styled(AccordionItemButton)`
  background: ${variables.linearBg} !important;
  padding: 12px;
  font-size: 26px;
  color: ${variables.text};
  cursor: pointer;
  display: flex;
  gap: 10px;
  align-items: center;
  justify-content: space-between;
  &[aria-expanded="true"] ${StyledChevron} {
    transform: rotate(-180deg);
  }
`;

const AccordionHeading = styled.div`
  font-size: 26px;
  color: ${variables.gray};
`;

const AccordionSubHeading = styled.div`
  font-size: 19px;
  margin-top: 10px;
  color: ${variables.white};
`;

const StyledAccordionItemPanel = styled(AccordionItemPanel)`
  background: ${variables.mainBg} !important;
  padding: 17px;
  &:not(:last-child) {
    border-bottom: 2px solid ${variables.borderBlue};
  }
`;

const AccordionItemNameContainer = styled.div`
  display: flex;
  align-items: center;
  gap: 8px;
`;

const TopicHeading = styled.h4`
  font-size: 24px;
  color: ${variables.text};
`;

const TopicSubHeading = styled.div`
  font-size: 17px;
  color: ${variables.white};
  margin-top: 7px;
`;
