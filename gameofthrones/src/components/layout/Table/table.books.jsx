import React, { useContext } from "react";
import { FetchDataContext } from "../../context/fetchDataContext";

function BooksTable() {
  const context = useContext(FetchDataContext);

  const { books } = context;

  return (
    <table style={{ border: "1px" }}>
      <thead>
        <tr>
          <td>Name</td>
          <td>ISBN</td>
          <td>Authors</td>
          <td>Number of pages</td>
          <td>Country</td>
        </tr>
      </thead>

      {books.map((book) => {
        return (
          <tbody>
            <tr>
              <td key={book.url} name={book.name}>
                {book.name}
              </td>
              <td key={book.isbn} name={book.isbn}>
                {book.isbn}
              </td>
              <td key={book.authors} name={book.authors}>
                {book.authors}
              </td>
              <td key={book.numberOfPages} name={book.numberOfPages}>
                {book.numberOfPages}
              </td>
              <td key={book.country} name={book.country}>
                {book.country}
              </td>
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

export default BooksTable;
