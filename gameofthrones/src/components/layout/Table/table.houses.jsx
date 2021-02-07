import React, { useContext } from "react";
import { FetchDataContext } from "../../context/fetchDataContext";

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
      {houses.map((house) => {
        return (
          <tbody>
            <tr>
              {house.name !== "" ? (
                <td key={house.url} name={house.name}>
                  {house.name}
                </td>
              ) : null}
              {house.name !== "" ? (
                <td key={house.region} name={house.region}>
                  {house.region}
                </td>
              ) : null}
              {house.name !== "" ? (
                <td key={house.words} name={house.words}>
                  {house.words}
                </td>
              ) : null}
              {house.name !== "" ? (
                <td key={house.title} name={house.title}>
                  {house.title}
                </td>
              ) : null}
            </tr>
          </tbody>
        );
      })}
      <tfoot>
        <tr></tr>
      </tfoot>
    </table>
  );
}

export default HousesTable;
