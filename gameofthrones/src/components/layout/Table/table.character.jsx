import React, { useContext } from 'react'
import { NavLink } from 'react-router-dom';
import { FetchDataContext } from '../../context/fetchDataContext'

function CharacterTable() {
    const context = useContext(FetchDataContext);

    const { characters } = context;
    return (
        <table>
            <thead>
                <tr>
                    <td>Name</td>
                    <td>Gender</td>
                    <td>Culture</td>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>
                        <ul>
                            {characters.map((character, i) => {
                                if (character.name === "") {
                                    return (
                                        <li key={character.url} name={character.name} >
                                            N/A
                                        </li>
                                    )
                                } else {
                                    return <li key={character.url} name={character.name} >{character.name}</li>
                                }
                            })}
                        </ul>
                    </td>

                    <td>
                        <ul>
                            {characters.map((character) => {
                                if (character.gender === "") {
                                    console.log('')
                                    return <li key={character.url} name={character.gender}>N/A</li>
                                } else {
                                    return <li key={character.url} name={character.gender} >{character.gender}</li>
                                }
                            })}
                        </ul>
                    </td>
                    <td>
                        <ul>
                            {characters.map((character) => {
                                if (character.culture === "") {
                                    return <li key={character.url} name={character.culture} >N/A</li>
                                } else {
                                    return <li key={character.url} name={character.culture} >{character.culture}</li>
                                }
                            })}
                        </ul>
                    </td>
                </tr>
            </tbody>
            <tfoot>
                <tr></tr>
            </tfoot>
        </table>
    )
}

export default CharacterTable;
