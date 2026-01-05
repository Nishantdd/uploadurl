import { SERVER_ADDRESS } from 'astro:env/client';
import React, { useState, useCallback } from 'react';

function FileUploader({ token }: { token: string }) {
    const [isDragging, setIsDragging] = useState(false);
    const [isUploading, setIsUploading] = useState(false);
    const [uploadError, setUploadError] = useState("");
    const [currentFileName, setCurrentFileName] = useState("");

    const onDragOver = useCallback((event: React.DragEvent) => {
        event.preventDefault();
        setIsDragging(true);
    }, []);

    const onDragLeave = useCallback(() => {
        setIsDragging(false);
    }, []);

    const onDrop = useCallback(async (event: React.DragEvent) => {
        event.preventDefault();
        setIsDragging(false);
        const files = event.dataTransfer.files;
        if(!files || !files[0]) setUploadError("Please provide an appropriate file") 
        else await handleFileUpload(files)
    }, []);

    const handleFileSelect = async (event: React.ChangeEvent<HTMLInputElement>) => {
        const files = event.target.files;
        if(!files || !files[0]) setUploadError("Please provide an appropriate file") 
        else await handleFileUpload(files)
    };

    const handleFileUpload = async (files: FileList) => {
        setIsUploading(true);
        setUploadError("");
        setCurrentFileName(files[0].name);
        
        let errorFlag = false;
        const formData = new FormData();
        formData.append("file", files[0]);
        
        await fetch(`${SERVER_ADDRESS}/api/file/upload`, {
            method: "POST",
            headers: { "Authorization": token },
            body: formData
        })
            .then(res => res.json())
            .then(res => { if(res.error) throw new Error(res.error) })
            .catch(_ => {
                errorFlag = true;
                setUploadError("Failed to upload the file. Please try again later.");
            })
            .finally(() => {
                if(errorFlag) setIsUploading(false);
                else window.location.reload();
            })
    }

    return (
        <div className="bg-gray-dark p-6 rounded-xl shadow-lg">
            <h2 className="text-2xl font-semibold mb-4 special-text">Upload Files</h2>
            {isUploading ? (
                <div className="border-2 border-dashed border-normal/20 rounded-xl p-8 text-center min-h-[200px] flex flex-col items-center justify-center">
                    <p className="text-normal/70 mb-4">{currentFileName}</p>
                    <div className="w-12 h-12 relative">
                        <div className="absolute inset-0 border-4 border-special/20 rounded-full"></div>
                        <div className="absolute inset-0 border-4 border-special rounded-full animate-spin border-t-transparent"></div>
                    </div>
                    <p className="text-normal/70 mt-4">Uploading...</p>
                </div>
            ) : (
                <div
                    className={`group border-2 border-dashed border-normal/20 rounded-xl p-8 text-center hover:border-special transition-colors cursor-pointer ${isDragging ? 'border-special' : ''}`}
                    onDragOver={onDragOver}
                    onDragLeave={onDragLeave}
                    onDrop={onDrop}
                    onClick={() => document.getElementById('fileInput')?.click()}
                >
                    <div className="text-normal/70">
                        <p className="mb-2">Drag and drop files here</p>
                        <p className="text-sm">or</p>
                        <label className="mt-2 px-4 py-2 font-medium group-hover:text-special rounded-md transition-colors cursor-pointer">
                            Browse Files
                        </label>
                        <input
                            type="file"
                            id="fileInput"
                            style={{ display: 'none' }}
                            onChange={handleFileSelect}
                        />
                    </div>
                </div>
            )}
            {uploadError && <p className="text-red-500 mt-2 text-sm">{uploadError}</p>}
        </div>
    );
};

export default FileUploader