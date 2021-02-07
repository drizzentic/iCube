import React, { useContext } from 'react'
import { FetchDataContext } from '../../context/fetchDataContext';

function HousesTable() {
    const context = useContext(FetchDataContext);

    const { houses } = context;

    return (
        <table>
            <thead>
                <tr>
                    <td>Name</td>
                    <td>Region</td>
                    <td>Words</td>
                    <td>Titles</td>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>
                        {houses.map(house => {
                            return (
                                <li key={house.url} name={house.name}>{house.name}</li>
                            )
                        })}
                    </td>
                    <td>
                        {houses.map(house => {
                            return (
                                <li key={house.url} name={house.region}>{house.region}</li>
                            )
                        })}
                    </td>
                    <td>
                        {houses.map(house => {
                            return (
                                <li key={house.url} name={house.words}>{house.words}</li>
                            )
                        })}
                    </td>
                    <td>
                        {houses.map(house => {
                            if (house.titles === "") {
                                return (
                                    <li key={house.url} name={house.titles}>{house.titles['']}</li>
                                )
                            } else {
                                return (
                                    <li key={house.url} name={house.titles}>{house.titles[0]}</li>
                                )
                            }
                        })}
                    </td>
                </tr>
            </tbody>
            <tfoot>
                <tr></tr>
            </tfoot>
        </table>
    )
}

export default HousesTable;
