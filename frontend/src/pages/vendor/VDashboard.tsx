import { VendorLayout } from '@/src/components/layout/VendorLayout'
import React from 'react'
import { Outlet } from 'react-router-dom'

const VDashboard = () => {
  return (
    <VendorLayout>
        <Outlet />
    </VendorLayout>
  )
}

export default VDashboard