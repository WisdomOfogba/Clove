import { useOnboardingStore } from "../../../store";
import { Icon } from "@iconify/react";
import { useState } from "react";
import { motion, AnimatePresence } from "motion/react";
import { cn, uploadToCloudinary } from "../../../lib/utils";
// Import the Dojah Connect Web Component / Hook depending on your version:
import Dojah from "dojah-kyc-sdk-react";

export function Step3() {
  const { data, updateProof, setStep } = useOnboardingStore();
  const userData = {
    first_name: data.user?.first_name || '', //Optional
    last_name: data.user?.last_name || '', //Optional
    dob: data.user?.dob || '', //YYYY-MM-DD Optional
    residence_country: 'NG', //Optional
    email: data.user?.email || ''//optional
  };


   const govData = {
    nin: '',
    bvn: '',
    dl: '',
    mobile: '',
 
  };
  const metadata = {
    user_id: '121',
  };
  // Localized image tracking state
  const [brandImages, setBrandImages] = useState<string[]>(
    data.proof.brandImages || [],
  );
  const [brandLogo, setBrandLogo] = useState<string | null>(
    data.proof.brandLogo || null,
  );
  const [dojahVerified, setDojahVerified] = useState<boolean>(
    !!data.proof.identityReference,
  );
  const [dojahRef, setDojahRef] = useState<string | null>(
    data.proof.identityReference || null,
  );

  // Per-image localized loading trackers
  const [logoUploading, setLogoUploading] = useState<boolean>(false);
  const [brandImagesUploading, setBrandImagesUploading] = useState<boolean[]>(
    [],
  );
  const [geoStatus, setGeoStatus] = useState<
    "idle" | "loading" | "success" | "error"
  >("idle");
  const [showDojah, setShowDojah] = useState<boolean>(false);

  // Handler for Brand Logo
  const handleLogoChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    setLogoUploading(true);
    try {
      const url = await uploadToCloudinary(file);
      setBrandLogo(url);
    } catch (err) {
      console.error("Logo upload failed", err);
    } finally {
      setLogoUploading(false);
    }
  };

  // Handler for Restaurant Images
  const handleBrandImagesChange = async (
    e: React.ChangeEvent<HTMLInputElement>,
  ) => {
    const files = Array.from(e.target.files || []);
    if (files.length === 0) return;

    // Expand our loading array map to account for the incoming pipeline items
    const newUploadTrackers = new Array(files.length).fill(true);
    setBrandImagesUploading((prev) => [...prev, ...newUploadTrackers]);

    // Process items in parallel safely
    files.forEach(async (file, index) => {
      try {
        const url = await uploadToCloudinary(file);
        setBrandImages((prev) => [...prev, url]);
      } catch (err) {
        console.error("Image upload failed", err);
      } finally {
        // Toggle specific loading block off
        setBrandImagesUploading((prev) => {
          const next = [...prev];
          next[index] = false;
          return next;
        });
      }
    });
  };

  const removeBrandImage = (index: number) => {
    setBrandImages((prev) => prev.filter((_, i) => i !== index));
  };

  const getLocation = () => {
    setGeoStatus("loading");
    navigator.geolocation.getCurrentPosition(
      (pos) => {
        updateProof({
          location: { lat: pos.coords.latitude, lng: pos.coords.longitude },
        });
        setGeoStatus("success");
      },
      () => setGeoStatus("error"),
    );
  };

  /**
   * Dojah SDK Configuration and Response Hook Callback
   */
  const dojahConfig = {
    appID: import.meta.env.VITE_DOJAH_APP_ID, // Replace with your Live/Sandbox App ID
    publicKey: import.meta.env.VITE_DOJAH_PUBLIC_KEY, // Replace with your Public Key
    type: "liveness", // Set to your EasyOnboard Flow Configuration ID if needed
    config: {
      debug: false,
      pages: [
        { page: "liveness", config: { selfie: true } },
      ],
    },
  };

  const handleDojahResponse = (type: string, data: any) => {
    console.log("Dojah Action Event:", type, data);

    if (type === "success") {
      // The user successfully finished the ID & selfie verification lifecycle
      setDojahVerified(true);
      setDojahRef(data.reference); // Store reference for backend verification lookup
      setShowDojah(false);
    } else if (type === "close") {
      setShowDojah(false);
    } else if (type === "error") {
      console.error("Dojah Widget pipeline error:", data);
      setShowDojah(false);
    }
  };

  const handleNext = () => {
    updateProof({
      brandImages: brandImages,
      brandLogo: brandLogo,
      identityReference: dojahRef,
      // Retain or sync state structure contextually
    });
    setStep(4);
  };

  // Verification requirements validation check
  const isFormValid =
    brandLogo &&
    !logoUploading &&
    brandImages.length >= 2 &&
    dojahVerified &&
    data.proof.location;

  return (
    <div className="p-8 space-y-8 relative overflow-hidden">
      <div className="space-y-1">
        <h2 className="text-xl font-bold text-dark">
          Restaurant Asset & Identity Verification
        </h2>
        <p className="text-sm text-neutral">
          Submit real corporate brand data and verify your administrative
          account identity.
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className="space-y-6">
          {/* Section 1: Brand Logo (Single Upload) */}
          <div className="space-y-3">
            <label className="text-xs font-bold text-neutral uppercase tracking-wider">
              Restaurant Brand Logo
            </label>
            <div className="flex items-center gap-4">
              {brandLogo ? (
                <div className="w-20 h-20 rounded-xl border border-slate-100 overflow-hidden relative group">
                  <img
                    src={brandLogo}
                    className="w-full h-full object-cover"
                    alt="Logo"
                  />
                  <button
                    onClick={() => setBrandLogo(null)}
                    className="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
                  >
                    <Icon
                      icon="lucide:trash-2"
                      className="w-5 h-5 text-white"
                    />
                  </button>
                </div>
              ) : (
                <label
                  className={cn(
                    "w-20 h-20 rounded-xl border-2 border-dashed flex flex-col items-center justify-center cursor-pointer transition-all",
                    logoUploading
                      ? "bg-slate-50 border-slate-300"
                      : "border-slate-200 hover:border-primary hover:bg-primary/5 text-slate-400 hover:text-primary",
                  )}
                >
                  {logoUploading ? (
                    <Icon
                      icon="lucide:loader-2"
                      className="w-6 h-6 animate-spin text-primary"
                    />
                  ) : (
                    <>
                      <Icon icon="lucide:upload-cloud" className="w-6 h-6" />
                      <input
                        type="file"
                        accept="image/*"
                        className="hidden"
                        onChange={handleLogoChange}
                        disabled={logoUploading}
                      />
                    </>
                  )}
                </label>
              )}
              <div className="space-y-0.5">
                <p className="text-xs font-bold text-dark">Brand Identifier</p>
                <p className="text-[10px] text-neutral">
                  Square PNG or JPG format recommended.
                </p>
              </div>
            </div>
          </div>

          {/* Section 2: Restaurant Brand Images (Multiple Uploads) */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <label className="text-xs font-bold text-neutral uppercase tracking-wider">
                Restaurant Showroom/Brand Gallery
              </label>
              <span className="text-[10px] font-bold text-neutral/50">
                {brandImages.length} / Min 2
              </span>
            </div>
            <div className="flex flex-wrap gap-2">
              {brandImages.map((img, i) => (
                <div
                  key={i}
                  className="relative w-16 h-16 rounded-lg overflow-hidden group border border-slate-100"
                >
                  <img
                    src={img}
                    className="w-full h-full object-cover"
                    alt="Gallery item"
                  />
                  <button
                    onClick={() => removeBrandImage(i)}
                    className="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 flex items-center justify-center transition-opacity"
                  >
                    <Icon icon="lucide:x" className="w-4 h-4 text-white" />
                  </button>
                </div>
              ))}

              {/* Individual Async Processing Loading Blocks representation */}
              {brandImagesUploading.map(
                (loading, idx) =>
                  loading && (
                    <div
                      key={`loading-${idx}`}
                      className="w-16 h-16 rounded-lg border border-slate-200 bg-slate-50 flex items-center justify-center"
                    >
                      <Icon
                        icon="lucide:loader-2"
                        className="w-4 h-4 animate-spin text-primary"
                      />
                    </div>
                  ),
              )}

              <label className="w-16 h-16 rounded-lg border-2 border-dashed border-slate-200 flex items-center justify-center hover:border-primary hover:bg-primary/5 cursor-pointer transition-all text-slate-400 hover:text-primary">
                <Icon icon="lucide:plus" className="w-6 h-6" />
                <input
                  type="file"
                  accept="image/*"
                  multiple
                  className="hidden"
                  onChange={handleBrandImagesChange}
                />
              </label>
            </div>
          </div>

          {/* Section 3: Dojah User Identity Check Integration */}
          <div className="space-y-3 p-4 bg-primary/5 rounded-2xl border border-primary/10">
            <div className="flex items-center justify-between">
              <label className="text-xs font-bold text-primary uppercase tracking-wider flex items-center gap-2">
                <Icon icon="lucide:shield-check" className="w-3.5 h-3.5" />{" "}
                Identity & Verification Hub
              </label>
            </div>
            <div className="flex items-center gap-4">
              {dojahVerified ? (
                <div className="w-12 h-12 rounded-full bg-emerald-100 flex items-center justify-center text-emerald-600">
                  <Icon icon="lucide:check-circle-2" className="w-6 h-6" />
                </div>
              ) : (
                <button
                  onClick={() => setShowDojah(true)}
                  className="w-12 h-12 rounded-full bg-primary/10 flex items-center justify-center text-primary hover:bg-primary/20 transition-all"
                >
                  <Icon icon="lucide:fingerprint" className="w-6 h-6" />
                </button>
              )}
              <div className="flex-1 space-y-0.5">
                <p className="text-xs font-bold text-dark">
                  Government Identity Proofing
                </p>
                <p className="text-[10px] text-neutral">
                  {dojahVerified
                    ? `Completed! Ref: ${dojahRef?.slice(0, 12)}...`
                    : "Powered by Dojah for secure live biometric authorization."}
                </p>
              </div>
            </div>
          </div>
        </div>

        {/* Location Section */}
        <div className="space-y-6">
          <div className="space-y-3">
            <label className="text-xs font-bold text-neutral uppercase tracking-wider block">
              Live Geolocation
            </label>
            <div className="h-48 bg-slate-50 rounded-2xl flex flex-col items-center justify-center border border-slate-200 relative overflow-hidden">
              {data.proof.location ? (
                <div className="absolute inset-0 flex flex-col items-center justify-center bg-white">
                  <Icon
                    icon="lucide:map-pin"
                    className="w-8 h-8 text-primary animate-bounce mb-2"
                  />
                  <p className="text-[10px] font-bold text-dark">
                    LOCATION PINNED
                  </p>
                  <p className="text-[10px] text-neutral">
                    {data.proof.location.lat.toFixed(4)},{" "}
                    {data.proof.location.lng.toFixed(4)}
                  </p>
                  <button
                    onClick={getLocation}
                    className="mt-2 text-[10px] font-bold text-primary hover:underline"
                  >
                    Repin
                  </button>
                </div>
              ) : (
                <>
                  <Icon
                    icon="lucide:map-pin"
                    className="w-8 h-8 text-slate-300 mb-2"
                  />
                  <button
                    onClick={getLocation}
                    disabled={geoStatus === "loading"}
                    className="bg-white px-4 py-2 rounded-xl text-xs font-bold shadow-sm hover:shadow-md transition-all flex items-center gap-2"
                  >
                    {geoStatus === "loading"
                      ? "Fetching..."
                      : "Verify Location"}
                    <Icon icon="lucide:upload" className="w-3 h-3" />
                  </button>
                </>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* Form Action Controls Footer */}
      <div className="flex gap-4">
        <button
          type="button"
          onClick={() => setStep(2)}
          className="flex-1 bg-slate-50 text-neutral rounded-xl py-4 font-bold border border-slate-100 hover:bg-slate-100 transition-all flex items-center justify-center gap-2"
        >
          <Icon icon="lucide:arrow-left" className="w-5 h-5" /> Back
        </button>
        <button
          onClick={handleNext}
          disabled={!isFormValid}
          className="flex-[2] bg-dark text-white rounded-xl py-4 font-bold hover:opacity-90 transition-all flex items-center justify-center gap-2 shadow-lg disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Review & Submit
          <Icon icon="lucide:arrow-right" className="w-5 h-5" />
        </button>
      </div>

      {/* Dojah Native Gateway Modal Overlay Layer */}
      <AnimatePresence>
        {showDojah && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            className="fixed inset-0 z-[100] bg-black/60 backdrop-blur-sm flex items-center justify-center p-4"
          >
            <div className="bg-white w-full max-w-xl rounded-3xl p-6 relative shadow-2xl overflow-hidden min-h-[500px] flex flex-col justify-between">
              <button
                onClick={() => setShowDojah(false)}
                className="absolute top-4 right-4 w-8 h-8 rounded-full bg-slate-100 flex items-center justify-center hover:bg-slate-200 transition-colors z-10"
              >
                <Icon icon="lucide:x" className="w-4 h-4" />
              </button>

              <div className="flex-1 mt-6">
                <Dojah
                  appID={dojahConfig.appID}
                  publicKey={dojahConfig.publicKey}
                  type={dojahConfig.type}
                  config={dojahConfig.config}
                  response={handleDojahResponse}
                  userData={userData}
                  govData={govData}
                  metadata={metadata}
                />
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
