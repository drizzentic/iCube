import React, { useContext } from "react";
import { FetchDataContext } from "../../context/fetchDataContext";

function Details() {
  const context = useContext(FetchDataContext);
  const { characters } = context;

  return (
    <>
      {characters.map((character, index) => {
        return (
          <table>
            <tbody>
              <tr>
                {character.name !== "" ? (
                  <>
                    <td key={index} name={character.name}>
                      {character.name}
                    </td>
                    <td>{character.gender}</td>
                    {character.culture !== "" ? (
                      <td>{character.culture}</td>
                    ) : null}
                    {character.books.length > 0 ? (
                      <td>
                        <ul>
                          {character.books.map((book) => {
                            return <li>{book}</li>;
                          })}
                        </ul>
                      </td>
                    ) : null}
                    {character.tvSeries.length > 0 ? (
                      <td>
                        <ul>
                          {character.tvSeries.map((tvSerie) => {
                            return <li>{tvSerie}</li>;
                          })}
                        </ul>
                      </td>
                    ) : null}
                  </>
                ) : null}
              </tr>
            </tbody>
          </table>
        );
      })}
    </>
  );
}

export default Details;
