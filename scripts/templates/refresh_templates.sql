SELECT template_id, template_name, template_description, template_content 
FROM speech_templates
WHERE template_id in $1
AND template_insert_date > $2;