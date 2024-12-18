import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios';

function App() {
  // state post
  const [cpfCnpj, setCpfCnpj] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  // state put
  const [cpfCnpjIdEdit, setCpfCnpjIdEdit] = useState('');
  const [cpfCnpjEdit, setCpfCnpjEdit] = useState('');
  const [cpfCnpjCheck, setCpfCnpjCheck] = useState(false);
  const [errorEdit, setErrorEdit] = useState('');
  const [successEdit, setSuccessEdit] = useState('');

  // state delete
  const [cpfCnpjIdDelete, setCpfCnpjIdDelete] = useState('');
  const [errorDelete, setErrorDelete] = useState('');
  const [successDelete, setSuccessDelete] = useState('');

  // state get
  const [documentos, setDocumentos] = useState([]);

  // função para cadastrar documento
  const handlePostCpfCnpj = async (e) => {
    e.preventDefault();

    setError('');
    setSuccess('');

    const dataCpfCnpj = {
      number: cpfCnpj,
    }
    try {
      const responseWriter = await axios.post('http://localhost:8080/cpf-cnpj', dataCpfCnpj);
      setCpfCnpj('');
      setSuccess('Documento cadastrado com sucesso!');
    } catch (err) {
      setError('Erro ao cadastrar documento.');
      console.error(err);
    }
  };

  // função para editar documento
  const handlePutCpfCnpj = async (e) => {
    e.preventDefault();

    setSuccessEdit('');
    setErrorEdit('');

    const dataCpfCnpj = {
      number: cpfCnpjEdit,
      blocked: cpfCnpjCheck,
    }

    try {
      const responseWriter = await axios.put(`http://localhost:8080/cpf-cnpj/${cpfCnpjIdEdit}`, dataCpfCnpj);
      //affectedRows vai depender do backend. Se ele já retornar no back, nasta por responseWriter.data
      if (responseWriter.data.affectedRows > 0) {
        setCpfCnpjIdEdit('');
        setCpfCnpjEdit('');
        setSuccessEdit('Documento editado com sucesso!');
      } else {
        setErrorEdit('Erro ao editar documento.');
      }
    } catch (err) {
      setErrorEdit('Erro ao editar documento.');
      console.error(err);
    }
  };

  // função para remover documento
  const handleDeleteCpfCnpj = async (e) => {
    e.preventDefault();

    setSuccessDelete('');
    setErrorDelete('');

    try {
      const responseWriter = await axios.delete(`http://localhost:8080/cpf-cnpj/${cpfCnpjIdDelete}`);
      //affectedRows vai depender do backend. Se ele já retornar no back, nasta por responseWriter.data
      if (responseWriter.data.affectedRows > 0) {
        setCpfCnpjIdDelete('');
        setSuccessDelete('Documento removido com sucesso!');
        setErrorDelete('');
      } else {
        setErrorDelete('Erro ao remover documento.');
        setSuccessDelete('');

      }
    } catch (err) {
      setErrorDelete('Erro ao remover documento.');
      setSuccessDelete('');
      console.error(err);
    }
  };

  // função para pegar os documentos
  useEffect(() => {
    const fetchDataDocumentos = async () => {
      try {
        const responseWriter = await axios.get('http://localhost:8080/cpf-cnpj');
        setDocumentos(responseWriter.data);
      } catch (err) {
        console.error('Erro ao buscar documentos:', err);
      }
    };
    fetchDataDocumentos();
  }, []);

  return (
    <div>
      <form onSubmit={handlePostCpfCnpj}>
        <input type="text" placeholder="CPF/CNPJ" value={cpfCnpj} onChange={(e) => setCpfCnpj(e.target.value)} />
        {success && <p>{success}</p>}
        {error && <p>{error}</p>}
        <button type="submit">Enviar</button>
      </form>

      <form onSubmit={handlePutCpfCnpj}>
        <input type="text" placeholder="ID" value={cpfCnpjIdEdit} onChange={(e) => setCpfCnpjIdEdit(e.target.value)} />
        <input type="text" placeholder="CPF/CNPJ" value={cpfCnpjEdit} onChange={(e) => setCpfCnpjEdit(e.target.value)} />
        <input type="checkbox" checked={cpfCnpjCheck} onChange={(e) => setCpfCnpjCheck(e.target.checked)} />
        {successEdit && <p>{successEdit}</p>}
        {errorEdit && <p>{errorEdit}</p>}
        <button type="submit">Editar</button>
      </form>

      <form onSubmit={handleDeleteCpfCnpj}>
        <input type="text" placeholder="ID" value={cpfCnpjIdDelete} onChange={(e) => setCpfCnpjIdDelete(e.target.value)} />
        {successDelete && <p>{successDelete}</p>}
        {errorDelete && <p>{errorDelete}</p>}
        <button type="submit">Remover</button>
      </form>

      {Array.isArray(documentos) && documentos.map((documento) => (
        <ul key={documento.id}>
          <li>ID: {documento.id}</li>
          <li>CPF/CNPJ: {documento.number}</li>
          <li>Bloqueado: {documento.blocked ? 'TRUE' : 'FALSE'}</li>
        </ul>
      ))}
    </div>
  );
}

export default App;
