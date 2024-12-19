import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './styles/app.css'

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

  //search
  const [search, setSearch] = useState('');
  const [filterDocumento, setFilterDocumento] = useState('');

  const handleSearchDocumento = (e) => {
    const pesquisado = e.target.value.toLowerCase();
    setSearch(pesquisado);

    const results = documentos.filter((documento) =>
      documento.number.toLowerCase().includes(pesquisado)
    );
    setFilterDocumento(results);
  }

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
      is_blocked: cpfCnpjCheck,
    }

    try {
      const responseWriter = await axios.put(`http://localhost:8080/cpf-cnpj/${cpfCnpjIdEdit}`, dataCpfCnpj);
      if (responseWriter.status === 200 || responseWriter.status === 201) {
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
        setFilterDocumento(responseWriter.data)
      } catch (err) {
        console.error('Erro ao buscar documentos:', err);
      }
    };
    fetchDataDocumentos();
  }, []);

  return (
    <div className='container-geral'>
      <h1 className='titulo-geral'>Validação CPF/CNPJ</h1>
      <div className='container-form'>
        <form className='formulario' onSubmit={handlePostCpfCnpj}>
          <h1 className='titulo-form'>Cadastrar Documento</h1>
          <input className='inpt-form' type="text" placeholder="CPF/CNPJ" value={cpfCnpj} onChange={(e) => setCpfCnpj(e.target.value)} />
          {success && <p className='success'>{success}</p>}
          {error && <p className='error'>{error}</p>}
          <button type="submit" className='btn-form'>Cadastrar</button>
        </form>

        <form className='formulario' onSubmit={handlePutCpfCnpj}>
          <h1 className='titulo-form'>Editar Documento</h1>
          <input className='inpt-form' type="text" placeholder="ID" value={cpfCnpjIdEdit} onChange={(e) => setCpfCnpjIdEdit(e.target.value)} />
          <input className='inpt-form' type="text" placeholder="CPF/CNPJ" value={cpfCnpjEdit} onChange={(e) => setCpfCnpjEdit(e.target.value)} />
          <div className='container-check'>
            <label htmlFor='isBlocked' className='label-check'>Bloqueado:</label>
            <input type="checkbox" name='isBlocked' checked={cpfCnpjCheck} onChange={(e) => setCpfCnpjCheck(e.target.checked)} />
          </div>
          {successEdit && <p className='success'>{successEdit}</p>}
          {errorEdit && <p className='error'>{errorEdit}</p>}
          <button type="submit" className='btn-form'>Editar</button>
        </form>

        <form className='formulario' onSubmit={handleDeleteCpfCnpj}>
          <h1 className='titulo-form'>Remover Documento</h1>
          <input className='inpt-form' type="text" placeholder="ID" value={cpfCnpjIdDelete} onChange={(e) => setCpfCnpjIdDelete(e.target.value)} />
          {successDelete && <p>{successDelete}</p>}
          {errorDelete && <p>{errorDelete}</p>}
          <button type="submit" className='btn-form'>Remover</button>
        </form>
      </div>
      <div>
        <input type='text' className='inpt-form' placeholder='Número do documento' value={search} onChange={handleSearchDocumento} />
      </div>
      <div className='container-lista'>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>CPF / CNPJ</th>
              <th>Bloqueado</th>
            </tr>
          </thead>
          <tbody>
            {Array.isArray(filterDocumento) && filterDocumento.map((documento) => (
              <tr key={documento.id}>
                <td>{documento.id}</td>
                <td>{documento.number}</td>
                <td>{documento.is_blocked ? 'Sim' : 'Não'}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default App;
