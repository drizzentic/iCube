import React, { Component, createContext } from 'react';


export const FetchDataContext = createContext();


export class FetchDataContextProvider extends Component {
    state = {
        characters: [],
        books: [],
        houses: [],
        missing: ''
    }

    componentDidMount() {
        const url1 = `https://anapioficeandfire.com/api/characters/`;
        const url2 = `https://anapioficeandfire.com/api/books/`;
        const url3 = `https://anapioficeandfire.com/api/houses/`;

        fetch(url1)
            .then(res => res.json())
            .then(jsonData => {
                // console.log(jsonData);

                if (jsonData.jsonData === "") {
                    console.log('input name missing')
                    this.setState({
                        missing: 'N/A'
                    })
                } else {
                    this.setState({
                        characters: jsonData,
                    });
                    console.log(jsonData)
                }
            })
        fetch(url2)
            .then(res => res.json())
            .then(jsonData => {
                // console.log(jsonData);

                if (jsonData.name === "") {
                    console.log('input name missing')
                    this.setState({
                        missing: 'N/A'
                    })
                } else {
                    this.setState({
                        books: jsonData,
                    });
                    console.log(jsonData)
                }
            })
        fetch(url3)
            .then(res => res.json())
            .then(jsonData => {
                // console.log(jsonData);

                if (jsonData.name === "") {
                    console.log('input name missing')
                    this.setState({
                        missing: 'N/A'
                    })
                } else {
                    this.setState({
                        houses: jsonData
                    });
                    console.log(jsonData)
                }
            })
    }
    render() {
        return (
            <FetchDataContext.Provider value={{
                characters: this.state.characters,
                // fetchData: this.fetchData,
                books: this.state.books,
                houses: this.state.houses
            }}>
                {this.props.children}
            </FetchDataContext.Provider>
        )
    }
}

export default FetchDataContextProvider
