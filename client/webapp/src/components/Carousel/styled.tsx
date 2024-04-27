import styled from "styled-components";

export const Container = styled.div`
  margin: 10px;

  h1 {
    margin-top: 0;
  }
`;

export const ListContainer = styled.div`
  height: 300px;
  overflow-y: hidden;
  overflow-x: scroll;
  transition: 500ms;
  scroll-behavior: smooth;

  scrollbar-width: none;
  &::-webkit-scrollbar {
    display: none;
  }
`;

export const Card = styled.div`
  height: 100%;
  &:first-child {
    div {
      margin-left: 0;
    }
  }
  &:last-child {
    div {
      margin-right: 0;
    }
  }
`;

export const Item = styled.div`
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
`;

export const ItemTitle = styled.h2`
  color: #fff;
  text-transform: capitalize;
  font-family: sans-serif;
  cursor: pointer;
`;

export const Title = styled.h1`
  text-transform: capitalize;
  font-family: sans-serif;
`;

// v.scroll-container {
//     height: 300px;
//     overflow-y: hidden;
//     overflow-x: scroll;
//     transition: 500ms;
//     scroll-behavior: smooth;

//     scrollbar-width: none;
//     &::-webkit-scrollbar {
//         display: none;
//     }
// }

// .card {
//     height: 100%;
//     &:first-child {
//         .box {
//             margin-left: 0;
//         }
//     }
//     &:last-child {
//         .box {
//             margin-right: 0;
//         }
//     }
// }
