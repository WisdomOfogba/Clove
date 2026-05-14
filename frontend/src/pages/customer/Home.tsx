import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Link } from 'react-router-dom';
import { useRef } from 'react';
import { cn } from '../../lib/utils';

export function CustomerHome() {
  const containerRef = useRef<HTMLDivElement>(null);

  return (
    <div ref={containerRef} className="pb-24 bg-[#FAFAFA] font-sans">
      {/* 1. HERO - Dribbble Inspired UI */}
      <section className="relative pt-12 pb-24 px-6 sm:px-12 lg:px-8 bg-white rounded-b-[3rem] shadow-[0_20px_50px_rgba(0,0,0,0.02)] border-b border-slate-100">
        <div className="container mx-auto max-w-7xl">
           <div className="flex flex-col lg:flex-row items-center gap-12 lg:gap-8 min-h-[70vh]">
              {/* Left Content */}
              <div className="flex-1 space-y-8 relative z-10 text-center lg:text-left mt-10 lg:mt-0">
                 <motion.div 
                   initial={{ opacity: 0, y: 10 }}
                   animate={{ opacity: 1, y: 0 }}
                   className="inline-flex items-center gap-2 px-5 py-2.5 rounded-full bg-orange-50 text-primary font-bold text-sm tracking-wide shadow-sm"
                 >
                    <span className="w-2.5 h-2.5 rounded-full bg-primary animate-pulse" />
                    #1 Verified Food Delivery
                 </motion.div>
                 
                 <motion.h1 
                   initial={{ opacity: 0, y: 20 }}
                   animate={{ opacity: 1, y: 0 }}
                   transition={{ delay: 0.1 }}
                   className="text-6xl sm:text-7xl lg:text-[5.5rem] font-extrabold tracking-tight text-dark leading-[1.05]"
                 >
                   The fastest <br className="hidden md:block" /> Delivery In <br className="hidden md:block" />
                   <span className="text-primary relative inline-block">
                     Your City
                     <svg className="absolute -bottom-2 sm:-bottom-4 left-0 w-full text-accent/30" viewBox="0 0 100 20" preserveAspectRatio="none">
                       <path d="M0,15 Q50,0 100,15" fill="none" stroke="currentColor" strokeWidth="6" strokeLinecap="round" />
                     </svg>
                   </span>
                 </motion.h1>

                 <motion.p 
                   initial={{ opacity: 0, y: 20 }}
                   animate={{ opacity: 1, y: 0 }}
                   transition={{ delay: 0.2 }}
                   className="max-w-lg mx-auto lg:mx-0 text-slate-500 text-lg sm:text-xl font-medium leading-relaxed"
                 >
                    enjoy a fast and smooth food delivery at your doorstep. Every kitchen is verified using Clove AI to ensure you’re ordering from authentic brands, not ghost vendors.
                 </motion.p>

                 {/* Search / Action */}
                 <motion.div 
                   initial={{ opacity: 0, y: 20 }}
                   animate={{ opacity: 1, y: 0 }}
                   transition={{ delay: 0.3 }}
                   className="flex flex-col sm:flex-row items-center gap-3 w-full max-w-xl mx-auto lg:mx-0 mt-8 bg-white p-2 sm:p-2.5 rounded-full shadow-[0_20px_40px_rgba(0,0,0,0.06)] border border-slate-100"
                 >
                    <div className="flex-1 flex items-center gap-3 px-5 w-full h-12">
                       <Icon icon="solar:map-point-outline" className="w-6 h-6 text-slate-400" />
                       <input 
                          type="text" 
                          placeholder="What's your delivery address?" 
                          className="w-full text-base font-semibold text-dark outline-none bg-transparent placeholder:text-slate-400"
                       />
                    </div>
                    <Link to="/restaurants" className="w-full sm:w-auto bg-primary text-white px-8 py-3.5 sm:py-4 rounded-full font-bold shadow-lg shadow-primary/30 hover:opacity-90 transition-all active:scale-95 flex items-center justify-center gap-2 shrink-0">
                       Order Now
                       <Icon icon="solar:arrow-right-bold" />
                    </Link>
                 </motion.div>

                 <motion.div 
                   initial={{ opacity: 0 }}
                   animate={{ opacity: 1 }}
                   transition={{ delay: 0.6 }}
                   className="flex items-center justify-center lg:justify-start gap-4 mt-6 text-sm font-bold text-slate-500"
                 >
                   <div className="flex items-center gap-1.5"><Icon icon="solar:shield-check-bold" className="text-green-500 w-5 h-5"/> Authentic Brands Only</div>
                   <div className="flex items-center gap-1.5"><Icon icon="solar:shield-check-bold" className="text-green-500 w-5 h-5"/> Powered by Clove AI</div>
                 </motion.div>
              </div>

              {/* Right Image Canvas */}
              <div className="flex-1 relative w-full lg:h-[600px] flex items-center justify-center mt-16 lg:mt-0">
                 {/* Decorative background blob */}
                 <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[350px] sm:w-[500px] h-[350px] sm:h-[500px] bg-primary/10 rounded-full blur-[80px]" />
                 <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-64 h-64 border-[40px] border-accent/5 rounded-full" />
                 
                 {/* Main Dish Image */}
                 <motion.div 
                   initial={{ opacity: 0, scale: 0.7, rotate: -15 }}
                   animate={{ opacity: 1, scale: 1, rotate: 0 }}
                   transition={{ duration: 1, type: "spring", bounce: 0.4 }}
                   className="relative z-10 w-[300px] sm:w-[450px] lg:w-[500px]"
                 >
                    <img src="https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&q=80&w=800&h=800" alt="Delicious Bowl" className="w-full h-full object-cover rounded-full shadow-[0_40px_80px_rgba(0,0,0,0.15)] relative z-10" />
                    
                    {/* Floating Leaves / Dots */}
                    <motion.div animate={{ y: [-10, 10, -10], rotate: [0, 10, 0] }} transition={{ repeat: Infinity, duration: 4, ease: "easeInOut" }} className="absolute -top-10 left-10 w-8 h-8 rounded-full bg-green-400/20 backdrop-blur-md border border-white flex items-center justify-center z-0 text-lg">🥬</motion.div>
                    <motion.div animate={{ y: [10, -10, 10], rotate: [0, -10, 0] }} transition={{ repeat: Infinity, duration: 5, ease: "easeInOut" }} className="absolute bottom-10 -right-5 w-12 h-12 rounded-full bg-accent/10 backdrop-blur-md border border-white flex items-center justify-center z-20 text-2xl">🥑</motion.div>

                    {/* Floating Cards */}
                    <motion.div 
                       initial={{ opacity: 0, x: -30 }}
                       animate={{ opacity: 1, x: 0 }}
                       transition={{ delay: 0.8, type: "spring" }}
                       className="absolute flex items-center gap-3 -left-4 sm:-left-12 top-20 bg-white/90 backdrop-blur-lg p-3 sm:p-4 rounded-2xl shadow-[0_20px_40px_rgba(0,0,0,0.08)] border border-white z-20"
                    >
                       <div className="w-10 h-10 sm:w-12 sm:h-12 bg-green-100 rounded-full flex items-center justify-center text-green-600">
                          <Icon icon="solar:shield-check-bold" className="w-5 h-5 sm:w-6 sm:h-6" />
                       </div>
                       <div>
                          <p className="text-xs sm:text-sm font-bold text-slate-500">Trust Score</p>
                          <p className="text-sm sm:text-base font-extrabold text-dark">99.8%</p>
                          <p className="text-[9px] font-bold text-primary mt-0.5 uppercase tracking-wide">Powered by Clove AI</p>
                       </div>
                    </motion.div>

                    <motion.div 
                       initial={{ opacity: 0, y: 30 }}
                       animate={{ opacity: 1, y: 0 }}
                       transition={{ delay: 1, type: "spring" }}
                       className="absolute -right-4 sm:-right-8 bottom-24 bg-white/90 backdrop-blur-lg p-3 sm:p-4 rounded-2xl shadow-[0_20px_40px_rgba(0,0,0,0.08)] border border-white z-20"
                    >
                       <div className="flex -space-x-2">
                          <img src="https://i.pravatar.cc/100?img=1" className="w-8 h-8 sm:w-10 sm:h-10 border-2 border-white rounded-full bg-slate-100" alt="User" />
                          <img src="https://i.pravatar.cc/100?img=5" className="w-8 h-8 sm:w-10 sm:h-10 border-2 border-white rounded-full bg-slate-100" alt="User" />
                          <div className="w-8 h-8 sm:w-10 sm:h-10 border-2 border-white rounded-full bg-accent flex items-center justify-center text-white text-[10px] sm:text-xs font-bold">50k+</div>
                       </div>
                       <div className="mt-2">
                          <p className="text-xs sm:text-sm font-bold text-slate-500 mb-0.5">Happy Customers</p>
                          <div className="flex text-primary">
                             {[...Array(5)].map((_, i) => <Icon key={i} icon="solar:star-bold" className="w-3 h-3 sm:w-4 sm:h-4" />)}
                          </div>
                       </div>
                    </motion.div>
                 </motion.div>
              </div>
           </div>
        </div>
      </section>

      {/* 2. HOW IT WORKS - Clean 3 Columns */}
      <section className="container mx-auto px-6 sm:px-12 mt-24">
         <div className="text-center max-w-2xl mx-auto mb-16">
            <h2 className="text-primary font-bold tracking-widest uppercase text-sm mb-3">Workflow</h2>
            <h3 className="text-3xl sm:text-4xl font-extrabold text-dark tracking-tight">How we serve you</h3>
         </div>
         <div className="grid grid-cols-1 md:grid-cols-3 gap-10 lg:gap-16 max-w-5xl mx-auto">
            {[
               { icon: "solar:shop-bold", title: "Easy to order", desc: "Browse highly vetted local restaurants and select your favorite fast food." },
               { icon: "solar:shield-check-bold", title: "Clove AI Verified", desc: "Every vendor is verified via NIN, CAC, and Smile ID to prevent ghost vendors." },
               { icon: "solar:delivery-bold", title: "Fastest Delivery", desc: "We deliver your food securely, unlocking the bag only at your location." }
            ].map((step, i) => (
               <div key={i} className="flex flex-col items-center text-center group cursor-pointer">
                  <div className="w-24 h-24 sm:w-28 sm:h-28 rounded-full bg-white shadow-[0_10px_30px_rgba(0,0,0,0.06)] border border-slate-100 flex items-center justify-center mb-6 relative group-hover:-translate-y-2 transition-transform duration-300">
                     <Icon icon={step.icon} className="w-10 h-10 text-dark group-hover:text-primary transition-colors" />
                     <div className="absolute -bottom-2 lg:-right-4 w-8 h-8 rounded-full bg-primary text-white flex items-center justify-center font-bold text-sm border-2 border-[#FAFAFA]">
                        0{i+1}
                     </div>
                  </div>
                  <h4 className="text-xl font-extrabold text-dark mb-3">{step.title}</h4>
                  <p className="text-slate-500 font-medium text-sm leading-relaxed">{step.desc}</p>
               </div>
            ))}
         </div>
      </section>

      {/* 3. POPULAR RESTAURANTS - Clean Cards */}
      <section className="container mx-auto px-6 sm:px-12 mt-32 max-w-7xl">
         <div className="flex flex-col sm:flex-row justify-between items-start sm:items-end mb-12">
            <div>
               <h2 className="text-primary font-bold tracking-widest uppercase text-sm mb-3">Verified Selection</h2>
               <h3 className="text-3xl sm:text-4xl font-extrabold text-dark tracking-tight">Trusted Restaurants</h3>
            </div>
            <Link to="/restaurants" className="mt-4 sm:mt-0 flex items-center gap-2 text-dark font-bold hover:text-primary transition-colors hover:gap-3">
               Explore Menu <Icon icon="solar:arrow-right-bold" />
            </Link>
         </div>

         <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {[
               { name: "The Greens Salad Bar", score: 99.8, type: "Healthy • Ikoyi", time: "15 min", img: "https://images.unsplash.com/photo-1546069901-d5bfd2cbfb1f?auto=format&fit=crop&q=80", price: "$12.00" },
               { name: "Spice Route Kitchen", score: 98.4, type: "Local Indian • VI", time: "30 min", img: "https://images.unsplash.com/photo-1565557623262-b51c2513a641?auto=format&fit=crop&q=80", price: "$18.50" },
               { name: "Artisan Bakes", score: 99.1, type: "Pastries • Lekki", time: "10 min", img: "https://images.unsplash.com/photo-1509440159596-0249088772ff?auto=format&fit=crop&q=80", price: "$8.00" },
            ].map((restaurant, i) => (
               <Link to={`/restaurants/${i}`} key={i} className="group block bg-white rounded-[2rem] p-5 shadow-[0_10px_40px_rgba(0,0,0,0.03)] border border-slate-100 hover:shadow-[0_20px_50px_rgba(0,0,0,0.06)] hover:-translate-y-1 transition-all">
                  <div className="relative aspect-[4/3] rounded-2xl overflow-hidden mb-5">
                     <img src={restaurant.img} alt={restaurant.name} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
                     <div className="absolute top-3 left-3 bg-white/95 backdrop-blur-sm px-3 py-1 rounded-full flex items-center gap-1.5 shadow-sm">
                        <Icon icon="solar:shield-check-bold" className="text-green-500 w-4 h-4" />
                        <span className="text-xs font-bold text-dark">{restaurant.score}</span>
                     </div>
                     <button className="absolute top-3 right-3 w-8 h-8 bg-white/95 backdrop-blur-sm rounded-full flex items-center justify-center text-slate-400 hover:text-red-500 transition-colors shadow-sm">
                        <Icon icon="solar:heart-bold" />
                     </button>
                  </div>
                  <div className="px-1">
                     <div className="flex justify-between items-start mb-2">
                        <h4 className="font-extrabold text-lg text-dark group-hover:text-primary transition-colors">{restaurant.name}</h4>
                        <span className="font-bold text-base text-dark">{restaurant.price}</span>
                     </div>
                     <p className="text-slate-500 text-sm font-medium mb-4">{restaurant.type}</p>
                     
                     <div className="flex items-center gap-4 text-xs font-bold text-slate-500">
                        <div className="flex items-center gap-1.5 bg-slate-50 px-3 py-1.5 rounded-lg border border-slate-100">
                           <Icon icon="solar:star-bold" className="w-4 h-4 text-primary" />
                           4.8 Rating
                        </div>
                        <div className="flex items-center gap-1.5 bg-slate-50 px-3 py-1.5 rounded-lg border border-slate-100">
                           <Icon icon="solar:clock-circle-bold-duotone" className="w-4 h-4 text-slate-400" />
                           {restaurant.time}
                        </div>
                     </div>
                  </div>
               </Link>
            ))}
         </div>
      </section>

      {/* 4. CTA BANNER */}
      <section className="container mx-auto px-6 sm:px-12 mt-32 max-w-7xl">
        <div className="relative rounded-[2.5rem] sm:rounded-[3rem] bg-accent overflow-hidden px-8 py-16 sm:p-20 text-center shadow-2xl flex flex-col items-center">
          {/* Abstract background blobs */}
          <div className="absolute -top-24 left-0 w-96 h-96 bg-white/10 rounded-full blur-3xl" />
          <div className="absolute bottom-0 right-0 w-96 h-96 bg-primary/20 rounded-full blur-3xl" />

          <div className="relative z-10 max-w-2xl space-y-8">
            <h2 className="text-4xl sm:text-5xl font-extrabold tracking-tight text-white leading-tight">
               Hungry? We are open <br /> with 100% authentic brands.
            </h2>
            <p className="text-white/80 text-base sm:text-lg font-medium mx-auto max-w-lg leading-relaxed">
              Order now and enjoy your food with full guarantee that you're ordering from the right restaurant.
            </p>
            <div className="pt-2 flex justify-center">
               <Link to="/restaurants" className="bg-primary text-white px-10 py-4 sm:py-5 rounded-full font-bold text-base flex items-center justify-center gap-3 hover:bg-white hover:text-dark transition-colors active:scale-95 shadow-xl shadow-dark/10">
                 Order Now
                 <Icon icon="solar:arrow-right-bold" className="w-5 h-5" />
               </Link>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
