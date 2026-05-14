import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Icon } from '@iconify/react';

export function VendorLayout({ children }: { children: React.ReactNode }) {
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <div className="min-h-screen bg-slate-50 flex">
      {/* Sidebar - Optional */}
      <div className="w-64 bg-white border-r border-slate-200 hidden lg:block h-screen fixed">
         <div className="p-8 pb-4">
            <h1 className="font-black text-2xl tracking-tighter text-dark mb-10">CLOVE <span className="text-primary italic">VENDORS</span></h1>
            <nav className="space-y-2">
               {[
                 { name: 'Overview', icon: 'solar:pie-chart-2-bold-duotone', path: '/vendor/dashboard' },
                 { name: 'Menu', icon: 'solar:menu-dots-square-bold-duotone', path: '/vendor/meals' },
                 { name: 'Orders', icon: 'solar:bell-bold-duotone', path: '/vendor/orders' },
                 { name: 'Profile', icon: 'solar:user-bold-duotone', path: '/vendor/profile' },
               ].map(item => {
                 const isActive = currentPath === item.path;
                 return (
                   <Link to={item.path} key={item.name} className={`flex items-center gap-3 px-4 py-3 rounded-2xl text-sm font-bold transition-all ${isActive ? 'bg-dark text-white shadow-xl' : 'text-neutral hover:bg-slate-50'}`}>
                      <Icon icon={item.icon} className="w-5 h-5" />
                      {item.name}
                   </Link>
                 )
               })}
            </nav>
         </div>
      </div>

      {/* Main Content Area */}
      <div className="flex-1 lg:ml-64 flex flex-col">
          {/* Mobile Header logic here if needed */}
          <div className="p-4 lg:p-10 flex-1">
             <div className="max-w-5xl mx-auto">
               {children}
             </div>
          </div>
      </div>
    </div>
  );
}
