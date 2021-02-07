import React, { useContext } from "react";
import { Link } from "react-router-dom";
import { FetchDataContext } from "../../context/fetchDataContext";

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

      {characters.map((character, index) => {
        return (
          <tbody>
            <tr>
              {character.name === "" ? (
                <td key={character.url} name={character.name}>
                  N/A
                </td>
              ) : (
                <td>
                  {" "}
                  <Link
                    to={`/character/details?${character.name}`}
                    target={"_blank"}
                  >
                    {character.name}
                  </Link>
                </td>
              )}
              {character.gender === "" ? (
                <td key={index} name={character.gender}>
                  N/A
                </td>
              ) : (
                <td key={index} name={character.gender}>
                  {character.gender}
                </td>
              )}
              {character.culture.length <= 0 ? (
                <td key={index} name={character.culture}>
                  N/A
                </td>
              ) : (
                <td key={index} name={character.culture}>
                  {character.culture}
                </td>
              )}
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

export default CharacterTable;
