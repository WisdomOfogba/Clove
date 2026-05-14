import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Icon } from '@iconify/react';

export function AdminLayout({ children }: { children: React.ReactNode }) {
  const location = useLocation();
  const currentPath = location.pathname;

  return (
    <div className="min-h-screen bg-slate-50 flex">
      {/* Sidebar */}
      <div className="w-64 bg-dark text-white border-r border-dark hidden lg:block h-screen fixed">
         <div className="p-8 pb-4">
            <h1 className="font-black text-2xl tracking-tighter text-white mb-10">CLOVE <span className="text-primary italic">ADMIN</span></h1>
            <nav className="space-y-2">
               {[
                 { name: 'Dashboard', icon: 'solar:pie-chart-2-bold-duotone', path: '/admin/dashboard' },
                 { name: 'Vendors', icon: 'solar:shop-bold-duotone', path: '/admin/vendors' },
                 { name: 'Review Queue', icon: 'solar:eye-bold-duotone', path: '/admin/review-queue' },
                 { name: 'Analytics', icon: 'solar:graph-up-bold-duotone', path: '/admin/analytics' },
               ].map(item => {
                 const isActive = currentPath === item.path || (currentPath === '/admin' && item.path === '/admin/dashboard');
                 return (
                   <Link to={item.path} key={item.name} className={`flex items-center gap-3 px-4 py-3 rounded-2xl text-sm font-bold transition-all ${isActive ? 'bg-primary text-dark shadow-xl' : 'text-slate-400 hover:text-white hover:bg-white/10'}`}>
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
          <div className="p-4 lg:p-10 flex-1">
             <div className="max-w-6xl mx-auto">
               {children}
             </div>
          </div>
      </div>
    </div>
  );
}
