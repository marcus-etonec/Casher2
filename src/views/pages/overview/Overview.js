import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import axios from 'axios'
import {
  CButton,
  CCard,
  CCardBody,
  CCardGroup,
  CCol,
  CContainer,
  CForm,
  CInput,
  CInputGroup,
  CInputGroupPrepend,
  CInputGroupText,
  CBadge,
  CCardHeader,
  CDataTable,
  CRow,
  CCardFooter,
  CLink,
  CWidgetProgress,
  CWidgetIcon,
  CWidgetProgressIcon,
  CWidgetSimple,
  CProgress
} from '@coreui/react'

import WidgetsBrand from '../../widgets/WidgetsBrand'
import WidgetsDropdown from '../../widgets/WidgetsDropdown'

import ChartLineSimple from '../../charts/ChartLineSimple'
import ChartBarSimple from '../../charts/ChartBarSimple'

import CIcon from '@coreui/icons-react'
import usersData from '../../users/UsersData'
const getBadge = status => {
  switch (status) {
    case 'Active': return 'success'
    case 'Inactive': return 'secondary'
    case 'Pending': return 'warning'
    case 'Banned': return 'danger'
    default: return 'primary'
  }
}
const fields = ['name','registered', 'role', 'status']


const Overview = () => {
  const [transactions, setTransactions] = useState([]);

  useEffect(() => {
      axios.get('http://localhost:8080/')
        .then((response) => {
          setTransactions(response.data.Data)
        });
  },[]);

  return (
    <div className="c-app c-default-layout flex-row align-items-center">
      <CContainer>
        <CRow>
          <CCol xs="12" sm="6" lg="3">
           <CWidgetProgressIcon
              header="87.500"
              text="Visitors"
              color="gradient-success"
              inverse
            >
              <CIcon name="cil-people" height="36"/>
            </CWidgetProgressIcon>
            </CCol>
            <CCol xs="12" sm="6" lg="3">
            <CWidgetProgressIcon
              header="87.500"
              text="Visitors"
              color="gradient-info"
              inverse
            >
              <CIcon name="cil-people" height="36"/>
            </CWidgetProgressIcon>  
          </CCol>
          <CCol xs="12" sm="6" lg="3">
            <CWidgetProgressIcon
              header="87.500"
              text="Visitors"
              color="gradient-warning"
              inverse
            >
              <CIcon name="cil-people" height="36"/>
            </CWidgetProgressIcon>  
          </CCol>
        </CRow>
        <CRow>
        <CCol>
          <CCard>
            <CCardHeader>
              Combined All dark Table
            </CCardHeader>
            <CCardBody>
            <CDataTable
              items={transactions}
              fields={fields}
              dark
              hover
              striped
              bordered
              size="sm"
              itemsPerPage={10}
              pagination
              scopedSlots = {{
                'status':
                  (item)=>(
                    <td>
                      <CBadge color={getBadge(item.status)}>
                        {item.status}
                      </CBadge>
                    </td>
                  )
              }}
            />
            </CCardBody>
          </CCard>
        </CCol>
      </CRow>
      </CContainer>
    </div>
  )
}

export default Overview
