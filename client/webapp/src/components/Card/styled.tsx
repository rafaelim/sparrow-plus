import styled from "styled-components";

export const Item = styled.div<{ active: boolean }>`
  position: relative;
  cursor: pointer;

  display: block;
  width: 250px;
  height: 300px;
  margin: 8px;
  cursor: pointer;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  > .overlay {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.7);
    padding: 20px;
    transition: all 0.5s ease;
    opacity: 0;
    pointer-events: none;
  }

  &:hover > .overlay {
    opacity: 1;
    pointer-events: all;
  }

  > .overlay .title {
    color: #e9e9e9;
    font-size: 16px;
    text-transform: capitalize;
    font-family: sans-serif;
    cursor: pointer;
  }

  > .overlay .title:hover {
    color: #fff;
  }

  ${({ active }) =>
    active &&
    `> .overlay {
    opacity: 1;

    pointer-events: all;}
  > .overlay .title {
    color: #fff;
  }
  `}
`;

export const ItemTitle = styled.h2`
  color: #fff;
  text-transform: capitalize;
  font-family: sans-serif;
  cursor: pointer;
`;
