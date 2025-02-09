export type urlResponse = {
    id: number,
    created_at: string,
    updated_at: string,
    original_url: string,
    short_url: string,
    user_id: number,
    slug: string
}

export type fileResponse = {
    id: number,
    created_at: string,
    updated_at: string,
    filename: string,
    filehash: string,
    file_type: string,
    file_size: number,
    location: string,
    user_id: number
}