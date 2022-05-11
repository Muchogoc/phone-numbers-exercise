import React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";

function App() {
  const [records, setRecords] = React.useState([]);
  const [isLoading, setIsLoading] = React.useState(true);

  React.useEffect(() => {
    var url = "http://localhost:8080/api/v1/customers";
    var headers = {};

    fetch(url, {
      method: "GET",
      mode: "cors",
      headers: headers,
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setRecords(data);
        setIsLoading(false);
      })
      .catch((error) => console.log(error));
  }, []);

  if (isLoading) return <p>Loading phone numbers</p>;

  return (
    <TableContainer component={Paper}>
      <Table stickyHeader>
        <TableHead>
          <TableCell>Name</TableCell>
          <TableCell>Country</TableCell>
          <TableCell>Phone Number</TableCell>
          <TableCell>State</TableCell>
        </TableHead>
        <TableBody>
          {records.map((record) => (
            <TableRow key={record.id}>
              <TableCell>{record.name}</TableCell>
              <TableCell>{record.country}</TableCell>
              <TableCell>{record.phone}</TableCell>
              <TableCell>{record.valid ? "valid" : "not valid"}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default App;
