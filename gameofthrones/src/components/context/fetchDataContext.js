import React, { Component, createContext } from "react";

export const FetchDataContext = createContext();

export class FetchDataContextProvider extends Component {
  state = {
    characters: [],
    books: [],
    houses: [],
  };

  componentDidMount() {
    const url1 = `https://anapioficeandfire.com/api/characters/`;
    const url2 = `https://anapioficeandfire.com/api/books/`;
    const url3 = `https://anapioficeandfire.com/api/houses/`;
    fetch(url1)
      .then((res) => res.json())
      .then((jsonData) => {
        if (jsonData.jsonData === "") {
          return;
        } else {
          this.setState({
            characters: jsonData,
          });
        }
      });
    fetch(url2)
      .then((res) => res.json())
      .then((jsonData) => {
        if (jsonData.name === "") {
          return;
        } else {
          this.setState({
            books: jsonData,
          });
        }
      });
    fetch(url3)
      .then((res) => res.json())
      .then((jsonData) => {
        if (jsonData.name === "") {
          console.log("input name missing");
        } else {
          this.setState({
            houses: jsonData,
          });
        }
      });
  }
  render() {
    return (
      <FetchDataContext.Provider
        value={{
          characters: this.state.characters,
          books: this.state.books,
          houses: this.state.houses,
        }}
      >
        {this.props.children}
      </FetchDataContext.Provider>
    );
  }
}

export default FetchDataContextProvider;
